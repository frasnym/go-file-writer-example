package asynchronousio

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	filewriter "github.com/frasnym/go-file-writer-example/file_writer"
)

type AsynchronousIOFileWriter struct {
	fileWriter    filewriter.FileWriter
	filename      string
	totalLines    int
	maxGoRoutines int
}

func NewAsynchronousIOFileWriter(totalLines int, filename string, fileWriter filewriter.FileWriter) *AsynchronousIOFileWriter {
	// Get the number of available CPU cores
	maxGoRoutines := runtime.GOMAXPROCS(0)

	return &AsynchronousIOFileWriter{
		totalLines:    totalLines,
		filename:      filename,
		fileWriter:    fileWriter,
		maxGoRoutines: maxGoRoutines,
	}
}

func (w *AsynchronousIOFileWriter) Write() error {
	// Create the output file
	file, err := w.fileWriter.CreateFile(w.filename)
	if err != nil {
		return err
	}
	defer w.fileWriter.FileClose(file)

	// Calculate the number of lines to be written by each worker
	linesPerTask := w.totalLines / w.maxGoRoutines

	var wg sync.WaitGroup
	errCh := make(chan error, w.maxGoRoutines)

	for i := 0; i < w.maxGoRoutines; i++ {
		wg.Add(1)
		go w.worker(i, file, &wg, linesPerTask, errCh)
	}

	// Close the error channel when all workers are done
	go func() {
		wg.Wait()
		close(errCh)
	}()

	// Collect and handle errors
	for err := range errCh {
		if err != nil {
			return err
		}
	}

	return nil
}

func (w *AsynchronousIOFileWriter) worker(id int, file *os.File, wg *sync.WaitGroup, linesPerTask int, errCh chan error) {
	defer wg.Done()
	startLine := id * linesPerTask
	endLine := startLine + linesPerTask

	for i := startLine; i < endLine; i++ {
		data := fmt.Sprintf("This is a line of data %d.\n", i)

		_, err := file.WriteString(data)
		if err != nil {
			errCh <- err
			return
		}
	}
}
