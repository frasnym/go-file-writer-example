package main

import (
	"flag"
	"fmt"
	"os"
)

// Function declarations for ease of unit testing.
var (
	// createFile is a function that creates a file with the given name.
	createFile = func(name string) (*os.File, error) {
		return os.Create(name)
	}

	// fileClose is a function that closes a file.
	fileClose = func(file *os.File) error {
		return file.Close()
	}

	// fileWriteString is a function that writes a string to a file and returns the number of bytes written.
	fileWriteString = func(file *os.File, s string) (n int, err error) {
		return file.WriteString(s)
	}
)

func main() {
	var lines int
	var filename string

	// Define command-line flags
	flag.IntVar(&lines, "lines", 100, "Total lines")
	flag.StringVar(&filename, "filename", "output/sequential_writing/one_hundred.txt", "File location")

	// Parse the command-line arguments
	flag.Parse()

	// Call writeToFile to create and write data to a file
	if err := writeToFile(lines, filename); err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File written successfully.")
}

// writeToFile writes the specified number of lines to the given file.
func writeToFile(lines int, filename string) (err error) {
	file, err := createFile(filename)
	if err != nil {
		return
	}
	defer fileClose(file)

	for i := 0; i < lines; i++ {
		_, err = fileWriteString(file, fmt.Sprintf("This is a line of data %d.\n", i))
		if err != nil {
			return
		}
	}

	return
}
