package parallelchunk

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	filewriter "github.com/frasnym/go-file-writer-example/file_writer"
)

type ParallelChunkFileWriter struct {
	fileWriter    filewriter.FileWriter
	filename      string
	totalLines    int
	maxGoRoutines int
}

func NewParallelChunkFileWriter(totalLines int, filename string, fileWriter filewriter.FileWriter) *ParallelChunkFileWriter {
	// Get the number of available CPU cores
	maxGoRoutines := runtime.GOMAXPROCS(0)

	return &ParallelChunkFileWriter{
		totalLines:    totalLines,
		filename:      filename,
		fileWriter:    fileWriter,
		maxGoRoutines: maxGoRoutines,
	}
}

func (w *ParallelChunkFileWriter) Write() error {
	// Create the output file
	file, err := w.fileWriter.CreateFile(w.filename)
	if err != nil {
		return err
	}
	w.fileWriter.FileClose(file)

	chunkSize := w.totalLines / w.maxGoRoutines
	var wg sync.WaitGroup

	for i := 0; i < w.maxGoRoutines; i++ {
		wg.Add(1)
		startLine := i * chunkSize
		endLine := startLine + chunkSize

		go w.writeChunkToFile(startLine, endLine, &wg)
	}

	wg.Wait()

	return nil
}

func (w *ParallelChunkFileWriter) writeChunkToFile(startLine, endLine int, wg *sync.WaitGroup) (err error) {
	file, err := w.fileWriter.OpenFile(w.filename, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return
	}
	defer w.fileWriter.FileClose(file)

	writer := w.fileWriter.NewBufferedWriter(file)

	for i := startLine; i < endLine; i++ {
		data := fmt.Sprintf("This is a line of data %d.\n", i)

		_, err = w.fileWriter.BufferedWriteString(writer, data)
		if err != nil {
			return
		}
	}

	w.fileWriter.BufferedFlush(writer)

	wg.Done()

	return
}
