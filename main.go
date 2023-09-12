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

	// Parse command-line flags for specifying file writing parameters.
	lines, filename, method := parseCommandLineFlags()

	// Get the appropriate file writing processor based on the specified method.
	processor := getProcessor(method)
	if processor == nil {
		fmt.Println("Undefined method")
		return
	}

	// Process the file writing task using the selected processor.
	fileWriter := filewriter.NewFileWriter()
	if err := processor(lines, filename, fileWriter); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("File written successfully, Binomial took %s\n", time.Since(start))
}

// parseCommandLineFlags parses and returns command-line flag values for lines, filename, and method.
func parseCommandLineFlags() (lines int, filename, method string) {
	flag.IntVar(&lines, "lines", 100, "Total lines")
	flag.StringVar(&filename, "filename", "output/sequential_writing/one_hundred.txt", "File location")
	flag.StringVar(&method, "method", "buffered_io", "File writer method")
	flag.Parse()

	return
}

// getProcessor returns the appropriate file writing processor function based on the specified method.
func getProcessor(method string) func(int, string, filewriter.FileWriter) error {
	mapFileWriterMethod := map[string]func(int, string, filewriter.FileWriter) error{
		"sequential":    fileWriterSequential,
		"parallel":      fileWriterParallel,
		"parallelchunk": fileWriterParallelChunk,
	}

	return mapFileWriterMethod[method]
}

// fileWriterParallel implements file writing in parallel mode.
func fileWriterParallel(totalLines int, filename string, fileWriter filewriter.FileWriter) error {
	fw := parallel.NewParallelFileWriter(fileWriter)
	return fw.Write(totalLines, filename)
}

// fileWriterSequential implements file writing in sequential mode.
func fileWriterSequential(totalLines int, filename string, fileWriter filewriter.FileWriter) error {
	fw := sequential.NewSequentialFileWriter(fileWriter)
	return fw.Write(totalLines, filename)
}

// fileWriterParallelChunk implements file writing in parallel mode with chunking.
func fileWriterParallelChunk(totalLines int, filename string, fileWriter filewriter.FileWriter) error {
	fw := parallelchunk.NewParallelChunkFileWriter(fileWriter)
	return fw.Write(totalLines, filename)
}
