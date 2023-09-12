package filewriter

type Writer interface {
	Write(totalLines int, filename string) error
}
