package sequentialwriting

import (
	"fmt"

	filewriter "github.com/frasnym/go-file-writer-example/file_writer"
)

type SequentialWritingFileWriter struct {
	fileWriter filewriter.FileWriter
	filename   string
	totalLines int
}

func NewSequentialWritingFileWriter(totalLines int, filename string, fileWriter filewriter.FileWriter) *SequentialWritingFileWriter {
	return &SequentialWritingFileWriter{
		totalLines: totalLines,
		filename:   filename,
		fileWriter: fileWriter,
	}
}

func (w *SequentialWritingFileWriter) Write() (err error) {
	// Create the output file
	file, err := w.fileWriter.CreateFile(w.filename)
	if err != nil {
		return
	}
	defer w.fileWriter.FileClose(file)

	for i := 0; i < w.totalLines; i++ {
		_, err = w.fileWriter.FileWriteString(file, fmt.Sprintf("This is a line of data %d.\n", i))
		if err != nil {
			return
		}
	}

	return
}
