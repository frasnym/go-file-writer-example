package parallel

import (
	"fmt"
	"os"
	"testing"

	"github.com/frasnym/go-file-writer-example/filewriter"
	"go.uber.org/mock/gomock"
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
			fw := NewParallelFileWriter(fileWriter)

			// Reset the timer for each benchmark iteration.
			b.ResetTimer()

			for i := 0; i < b.N; i++ {
				// Run the file-writing function with the current number of lines.
				err := fw.Write(lines, tmpDir+"/benchmark.txt")
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

func TestParallelFileWriter_Write(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock for the FileWriter interface
	mockFileWriter := filewriter.NewMockFileWriter(ctrl)

	// Create an instance of ParallelFileWriter with the mock FileWriter
	pfw := NewParallelFileWriter(mockFileWriter)

	// Set up expectations for the mock FileWriter
	totalLines := 100 // Set your desired totalLines and filename
	filename := "test.txt"
	mockFileWriter.EXPECT().CreateFile(filename).Return(nil, nil)
	mockFileWriter.EXPECT().NewBufferedWriter(gomock.Any()).Return(nil).AnyTimes()
	mockFileWriter.EXPECT().BufferedWriteString(gomock.Any(), gomock.Any()).Return(0, nil).AnyTimes()
	mockFileWriter.EXPECT().BufferedFlush(gomock.Any()).Return(nil).AnyTimes()
	mockFileWriter.EXPECT().FileClose(gomock.Any()).Return(nil)

	// Run the Write function
	err := pfw.Write(totalLines, filename)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
}
