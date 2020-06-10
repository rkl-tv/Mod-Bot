package memory

import "io"

type WriterFactory interface {
	New(process interface{}, address uintptr) io.Writer
}
