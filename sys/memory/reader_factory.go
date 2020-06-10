package memory

import "io"

type ReaderFactory interface {
	New(process interface{}, address uintptr) io.Reader
}
