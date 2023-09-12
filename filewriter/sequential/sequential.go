package sequential

import (
	"fmt"

	"github.com/frasnym/go-file-writer-example/filewriter"
)

// SequentialFileWriter represents a writer for sequential file operations.
type sequentialFileWriter struct {
	fileWriter filewriter.FileWriter
}

// NewSequentialFileWriter creates a new instance of SequentialFileWriter.
func NewSequentialFileWriter(fileWriter filewriter.FileWriter) filewriter.Writer {
	return &sequentialFileWriter{
		fileWriter: fileWriter,
	}
}

// Write writes the specified number of lines to the file sequentially.
func (w *sequentialFileWriter) Write(totalLines int, filename string) error {
	// Create the output file
	file, err := w.fileWriter.CreateFile(filename)
	if err != nil {
		return err
	}
	defer w.fileWriter.FileClose(file)

	writer := w.fileWriter.NewBufferedWriter(file)
	for i := 0; i < totalLines; i++ {
		data := fmt.Sprintf("This is a line of data %d.\n", i)
		_, err := w.fileWriter.BufferedWriteString(writer, data)
		if err != nil {
			return err
		}
	}

	// Flush the buffer to ensure all data is written to the file.
	err = w.fileWriter.BufferedFlush(writer)
	if err != nil {
		return err
	}

	return nil
}
