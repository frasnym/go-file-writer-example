package filewriter

type Writer interface {
	Write() (err error)
}
