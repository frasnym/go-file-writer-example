package main

import (
	"flag"
	"fmt"

	filewriter "github.com/frasnym/go-file-writer-example/file_writer"
	"github.com/frasnym/go-file-writer-example/file_writer/parallel"
	"github.com/frasnym/go-file-writer-example/file_writer/sequential"
)

func main() {
	var (
		lines    int
		filename string
		method   string
	)

	// Define command-line flags
	flag.IntVar(&lines, "lines", 100, "Total lines")
	flag.StringVar(&filename, "filename", "output/sequential_writing/one_hundred.txt", "File location")
	flag.StringVar(&method, "method", "buffered_io", "File writer method")

	// Parse the command-line arguments
	flag.Parse()

	mapFileWriterMethod := map[string]func(totalLines int, filename string, fileWriter filewriter.FileWriter) error{
		"parallel":   fileWriterParallel,
		"sequential": fileWriterSequential,
	}

	processor := mapFileWriterMethod[method]
	if processor == nil {
		fmt.Println("Undefined method")
		return
	}

	// Call writeToFile to create and write data to a file
	fileWriter := filewriter.NewFileWriter()
	if err := processor(lines, filename, fileWriter); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File written successfully.")
}

func fileWriterParallel(totalLines int, filename string, fileWriter filewriter.FileWriter) error {
	fw := parallel.NewParallelFileWriter(totalLines, filename, fileWriter)
	return fw.Write()
}

func fileWriterSequential(totalLines int, filename string, fileWriter filewriter.FileWriter) error {
	fw := sequential.NewSequentialFileWriter(totalLines, filename, fileWriter)
	return fw.Write()
}
