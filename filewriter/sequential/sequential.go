package sequential

import (
	"fmt"

	filewriter "github.com/frasnym/go-file-writer-example/filewriter"
)

type SequentialFileWriter struct {
	fileWriter filewriter.FileWriter
	filename   string
	totalLines int
}

func NewSequentialFileWriter(totalLines int, filename string, fileWriter filewriter.FileWriter) *SequentialFileWriter {
	return &SequentialFileWriter{
		totalLines: totalLines,
		filename:   filename,
		fileWriter: fileWriter,
	}
}

func (w *SequentialFileWriter) Write() (err error) {
	// Create the output file
	file, err := w.fileWriter.CreateFile(w.filename)
	if err != nil {
		return
	}
	defer w.fileWriter.FileClose(file)

	writer := w.fileWriter.NewBufferedWriter(file)
	for i := 0; i < w.totalLines; i++ {
		_, err = w.fileWriter.BufferedWriteString(writer, fmt.Sprintf("This is a line of data %d.\n", i))
		if err != nil {
			return
		}
	}

	w.fileWriter.BufferedFlush(writer)

	return
}
