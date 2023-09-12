package main

import (
	"flag"
	"fmt"
	"time"

	"github.com/frasnym/go-file-writer-example/filewriter"
	"github.com/frasnym/go-file-writer-example/filewriter/parallel"
	"github.com/frasnym/go-file-writer-example/filewriter/parallelchunk"
	"github.com/frasnym/go-file-writer-example/filewriter/sequential"
)

func main() {
	start := time.Now()
	lines, filename, method := parseCommandLineFlags()

	fileWriter := filewriter.NewFileWriter()
	processor := getProcessor(method)

	if processor == nil {
		fmt.Println("Undefined method")
		return
	}

	if err := processFile(lines, filename, fileWriter, processor); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File written successfully, Binomial took %s\n", time.Since(start))
}

func parseCommandLineFlags() (int, string, string) {
	var lines int
	var filename, method string

	flag.IntVar(&lines, "lines", 100, "Total lines")
	flag.StringVar(&filename, "filename", "output/sequential_writing/one_hundred.txt", "File location")
	flag.StringVar(&method, "method", "buffered_io", "File writer method")
	flag.Parse()

	return lines, filename, method
}

func getProcessor(method string) func(int, string, filewriter.FileWriter) error {
	mapFileWriterMethod := map[string]func(int, string, filewriter.FileWriter) error{
		"sequential":    fileWriterSequential,
		"parallel":      fileWriterParallel,
		"parallelchunk": fileWriterParallelChunk,
	}

	return mapFileWriterMethod[method]
}

func processFile(lines int, filename string, fileWriter filewriter.FileWriter, processor func(int, string, filewriter.FileWriter) error) error {
	return processor(lines, filename, fileWriter)
}

func fileWriterParallel(totalLines int, filename string, fileWriter filewriter.FileWriter) error {
	fw := parallel.NewParallelFileWriter(totalLines, filename, fileWriter)
	return fw.Write()
}

func fileWriterSequential(totalLines int, filename string, fileWriter filewriter.FileWriter) error {
	fw := sequential.NewSequentialFileWriter(totalLines, filename, fileWriter)
	return fw.Write()
}

func fileWriterParallelChunk(totalLines int, filename string, fileWriter filewriter.FileWriter) error {
	fw := parallelchunk.NewParallelChunkFileWriter(totalLines, filename, fileWriter)
	return fw.Write()
}
