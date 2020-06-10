package windows

import (
	"ModBot/sys/process"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestMemWriterFactory_New(t *testing.T) {
	f := NewMemWriterFactory()

	w := f.New(process.DWORD(666), 0x123456)
	assert.IsType(t, &writer{}, w)
	assert.Implements(t, (*io.Writer)(nil), w)
}
