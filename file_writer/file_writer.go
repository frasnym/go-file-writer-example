package filewriter

import "os"

type FileWriter interface {
	CreateFile(name string) (*os.File, error)
	FileWriteString(file *os.File, s string) (n int, err error)
	FileWrite(file *os.File, b []byte) (n int, err error)
	FileClose(file *os.File) error
}

type fileWriter struct{}

func NewFileWriter() FileWriter {
	return &fileWriter{}
}

func (r *fileWriter) CreateFile(name string) (*os.File, error) {
	return os.Create(name)
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
