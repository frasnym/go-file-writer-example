package filewriter

import (
	"bufio"
	"io/fs"
	"os"
)

type FileWriter interface {
	CreateFile(name string) (*os.File, error)
	OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error)
	FileWriteString(file *os.File, s string) (n int, err error)
	FileWrite(file *os.File, b []byte) (n int, err error)
	FileClose(file *os.File) error
	NewBufferedWriter(file *os.File) *bufio.Writer
	BufferedWriteString(writer *bufio.Writer, s string) (n int, err error)
	BufferedFlush(writer *bufio.Writer) error
}

type fileWriter struct{}

func NewFileWriter() FileWriter {
	return &fileWriter{}
}

func (r *fileWriter) CreateFile(name string) (*os.File, error) {
	return os.Create(name)
}

func (*fileWriter) OpenFile(name string, flag int, perm fs.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (*fileWriter) FileWrite(file *os.File, b []byte) (n int, err error) {
	return file.Write(b)
}

func (r *fileWriter) FileWriteString(file *os.File, s string) (n int, err error) {
	return file.WriteString(s)
}

func (r *fileWriter) FileClose(file *os.File) error {
	return file.Close()
}

func (r *fileWriter) NewBufferedWriter(file *os.File) *bufio.Writer {
	return bufio.NewWriter(file)
}

func (r *fileWriter) BufferedWriteString(writer *bufio.Writer, s string) (n int, err error) {
	return writer.WriteString(s)
}

func (r *fileWriter) BufferedFlush(writer *bufio.Writer) error {
	return writer.Flush()
}
