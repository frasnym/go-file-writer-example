package parallelchunk

import (
	"fmt"
	"runtime"
	"sync"

	"github.com/frasnym/go-file-writer-example/filewriter"
)

type parallelChunkFileWriter struct {
	fileWriter    filewriter.FileWriter
	maxGoRoutines int
}

func NewParallelChunkFileWriter(fileWriter filewriter.FileWriter) filewriter.Writer {
	// Get the number of available CPU cores
	maxGoRoutines := runtime.GOMAXPROCS(0)

	return &parallelChunkFileWriter{
		fileWriter:    fileWriter,
		maxGoRoutines: maxGoRoutines,
	}
}

func (w *parallelChunkFileWriter) Write(totalLines int, filename string) error {
	chunkSize := totalLines / w.maxGoRoutines
	var wg sync.WaitGroup

	for i := 0; i < w.maxGoRoutines; i++ {
		wg.Add(1)
		startLine := i * chunkSize
		endLine := startLine + chunkSize

		go w.writeChunkToFile(startLine, endLine, filename, &wg)
	}

	wg.Wait()

	return nil
}

func (w *parallelChunkFileWriter) writeChunkToFile(startLine, endLine int, filename string, wg *sync.WaitGroup) (err error) {
	file, err := w.fileWriter.CreateFile(fmt.Sprint(filename, "_", startLine))
	if err != nil {
		return err
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
