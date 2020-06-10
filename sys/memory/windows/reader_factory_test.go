package windows

import (
	"ModBot/sys/process"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestMemReaderFactory_New(t *testing.T) {
	f := NewMemReaderFactory()

	r := f.New(process.DWORD(666), 0x123456)
	assert.IsType(t, &reader{}, r)
	assert.Implements(t, (*io.Reader)(nil), r)
}
