package sequential

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
			fw := NewSequentialFileWriter(fileWriter)

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

func TestSequentialFileWriter_Write(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Create a mock of the FileWriter interface.
	mockFileWriter := filewriter.NewMockFileWriter(ctrl)

	// Create a SequentialFileWriter with the mock FileWriter.
	writer := NewSequentialFileWriter(mockFileWriter)

	// Specify the expected parameters and behavior for the FileWriter methods.
	filename := "test.txt"
	totalLines := 5
	expectedData := "This is a line of data"

	// Expect calls to the FileWriter methods.
	mockFileWriter.EXPECT().CreateFile(filename).Return(nil, nil)
	mockFileWriter.EXPECT().NewBufferedWriter(gomock.Any()).Return(nil)
	for i := 0; i < totalLines; i++ {
		mockFileWriter.EXPECT().BufferedWriteString(gomock.Any(), fmt.Sprintf("%s %d.\n", expectedData, i)).Return(0, nil)
	}
	mockFileWriter.EXPECT().BufferedFlush(gomock.Any()).Return(nil)
	mockFileWriter.EXPECT().FileClose(gomock.Any()).Return(nil)

	// Call the Write method of SequentialFileWriter.
	err := writer.Write(totalLines, filename)

	// Check if there were any errors during the Write operation.
	if err != nil {
		t.Errorf("error should be nil, but got: %v", err)
	}
}
