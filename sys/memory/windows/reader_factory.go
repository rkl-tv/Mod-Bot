package windows

import (
	"ModBot/sys/memory"
	process2 "ModBot/sys/process"
	"io"
)

type MemReaderFactory struct {
}

func NewMemReaderFactory() memory.ReaderFactory {
	return &MemReaderFactory{}
}

func (f *MemReaderFactory) New(process interface{}, offset uintptr) io.Reader {
	pid, _ := process.(process2.DWORD) // TODO, ok so or handle panic?

	return newReader(pid, offset)
}
