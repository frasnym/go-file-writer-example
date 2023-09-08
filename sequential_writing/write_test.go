package main

import (
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

var errMock = errors.New("an error")

func BenchmarkWriteToFileExecutionTime(b *testing.B) {
	// Create a temporary directory for writing files.
	tmpDir, err := os.MkdirTemp("", "file-benchmark")
	if err != nil {
		b.Fatalf("Error creating temporary directory: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	// Specify the filename for writing.
	filename := tmpDir + "/benchmark.txt"

	b.ResetTimer() // Reset the timer to exclude setup time.

	for i := 0; i < b.N; i++ {
		// Record the start time.
		startTime := time.Now()

		// Run the file-writing function.
		err := writeToFile(10_000_000, filename)
		if err != nil {
			b.Fatalf("Error writing to file: %v", err)
		}

		// Calculate the execution time.
		elapsedTime := time.Since(startTime)
		// Log the execution time (you can also collect it for analysis).
		fmt.Printf("Execution Time: %s\n", elapsedTime)
	}
}

func TestMain(t *testing.T) {
	// Backup the original command-line arguments and stdout
	oldArgs := os.Args
	oldOutput := os.Stdout

	// Ensure that these are restored after the test
	defer func() {
		os.Args = oldArgs
		os.Stdout = oldOutput
	}()

	// Capture the output for testing
	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create a temporary flag set for testing
	fs := flag.NewFlagSet("test", flag.ExitOnError)
	var lines int
	var filename string
	fs.IntVar(&lines, "lines", 0, "")
	fs.StringVar(&filename, "filename", "", "")

	// Set the flag values for testing
	testFileName := "testFile.txt"
	os.Args = []string{"test", "--lines=5", fmt.Sprintf("--filename=%s", testFileName)}

	// Parse the command-line arguments
	fs.Parse(os.Args[1:])

	// Call the main function
	main()

	// Close the write end of the pipe and capture the output
	w.Close()
	out, _ := io.ReadAll(r)

	// Assertions
	expectedOutput := "File written successfully.\n"
	if string(out) != expectedOutput {
		t.Errorf("Expected output: %q, but got: %q", expectedOutput, string(out))
	}

	// Clean up the file
	err := os.Remove(testFileName)
	if err != nil {
		t.Errorf("Error cleaning up %s: %v", testFileName, err)
	}
}

func TestWriteToFile(t *testing.T) {
	testCases := []struct {
		name                string
		mockCreateFile      func(name string) (*os.File, error)
		mockFileClose       func(file *os.File) error
		mockFileWriteString func(file *os.File, s string) (n int, err error)
		wantError           bool
	}{
		{
			name: "happy path",
			mockCreateFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockFileClose: func(file *os.File) error {
				return nil
			},
			mockFileWriteString: func(file *os.File, s string) (n int, err error) {
				return 0, nil
			},
			wantError: false,
		},
		{
			name: "error create file",
			mockCreateFile: func(name string) (*os.File, error) {
				return nil, errMock
			},
			wantError: true,
		},
		{
			name: "error write file",
			mockCreateFile: func(name string) (*os.File, error) {
				return new(os.File), nil
			},
			mockFileClose: func(file *os.File) error {
				return nil
			},
			mockFileWriteString: func(file *os.File, s string) (n int, err error) {
				return 0, errMock
			},
			wantError: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			createFile = tc.mockCreateFile
			fileClose = tc.mockFileClose
			fileWriteString = tc.mockFileWriteString

			err := writeToFile(1, "filename")
			isError := err != nil

			if isError != tc.wantError {
				t.Errorf("Expected %t, but got %t", tc.wantError, isError)
			}
		})
	}
}
