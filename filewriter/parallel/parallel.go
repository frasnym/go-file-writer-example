package parallel

import (
	"fmt"
	"os"
	"runtime"
	"sync"

	filewriter "github.com/frasnym/go-file-writer-example/filewriter"
)

type ParallelFileWriter struct {
	fileWriter    filewriter.FileWriter
	maxGoRoutines int
}

func NewParallelFileWriter(fileWriter filewriter.FileWriter) *ParallelFileWriter {
	// Get the number of available CPU cores
	maxGoRoutines := runtime.GOMAXPROCS(0)

	return &ParallelFileWriter{
		fileWriter:    fileWriter,
		maxGoRoutines: maxGoRoutines,
	}
}

func (w *ParallelFileWriter) Write(totalLines int, filename string) error {
	// Create the output file
	file, err := w.fileWriter.CreateFile(filename)
	if err != nil {
		return err
	}
	defer w.fileWriter.FileClose(file)

	// Calculate the number of lines to be written by each worker
	linesPerTask := totalLines / w.maxGoRoutines

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

func (w *ParallelFileWriter) worker(id int, file *os.File, wg *sync.WaitGroup, linesPerTask int, errCh chan error) {
	defer wg.Done()
	startLine := id * linesPerTask
	endLine := startLine + linesPerTask

	writer := w.fileWriter.NewBufferedWriter(file)
	for i := startLine; i < endLine; i++ {
		data := fmt.Sprintf("This is a line of data %d.\n", i)

		_, err := w.fileWriter.BufferedWriteString(writer, data)
		if err != nil {
			errCh <- err
			return
		}
	}

	w.fileWriter.BufferedFlush(writer)
}
