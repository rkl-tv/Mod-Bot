package windows

import (
	"ModBot/sys/memory"
	process2 "ModBot/sys/process"
	"io"
)

type MemWriterFactory struct {
}

func NewMemWriterFactory() memory.WriterFactory {
	return &MemWriterFactory{}
}

func (f *MemWriterFactory) New(process interface{}, offset uintptr) io.Writer {
	pid, _ := process.(process2.DWORD) // TODO, ok so or handle panic?

	return newWriter(pid, offset)
}
