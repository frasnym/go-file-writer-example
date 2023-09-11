package sequentialwriting

import (
	"fmt"
	"os"
	"testing"

	filewriter "github.com/frasnym/go-file-writer-example/file_writer"
)

func BenchmarkWriteToFileExecutionTime(b *testing.B) {
	// Create a temporary directory for writing files.
	tmpDir, cleanup := createTempDir(b)
	defer cleanup()

	// Define an array of data sizes (number of lines) to benchmark.
	dataSizes := []int{10, 100, 1000, 10_000, 100_000, 1000_000, 10_000_000} // Example data sizes.

	for _, lines := range dataSizes {
		b.Run(fmt.Sprintf("Lines-%d", lines), func(b *testing.B) {
			// Create a FileWriter and an AsynchronousIOFileWriter for each benchmark iteration.
			fileWriter := filewriter.NewFileWriter()
			fw := NewSequentialWritingFileWriter(lines, tmpDir+"/benchmark.txt", fileWriter)

			// Reset the timer for each benchmark iteration.
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				// Run the file-writing function with the current number of lines.
				err := fw.Write()
				if err != nil {
					b.Fatalf("Error writing to file: %v", err)
				}
			}
		})
	}
}

// createTempDir creates a temporary directory and returns its path along with a cleanup function.
func createTempDir(b *testing.B) (string, func()) {
	b.Helper()

	tmpDir, err := os.MkdirTemp("", "file-benchmark")
	if err != nil {
		b.Fatalf("Error creating temporary directory: %v", err)
	}

	cleanup := func() {
		os.RemoveAll(tmpDir)
	}

	return tmpDir, cleanup
}
