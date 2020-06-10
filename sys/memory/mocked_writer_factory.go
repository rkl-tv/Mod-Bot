package memory

import "io"

type MockedWriterFactory struct {
	NewFunc func(process interface{}, address uintptr) io.Writer
}

func (f *MockedWriterFactory) New(process interface{}, address uintptr) io.Writer {
	if f.NewFunc != nil {
		return f.NewFunc(process, address)
	}

	return nil
}
