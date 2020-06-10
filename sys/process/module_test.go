package process

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewModule(t *testing.T) {
	path := "/foo/bar"
	bAddr := uintptr(1234567890987654321)

	m := NewModule(path, bAddr)
	assert.NotNil(t, m)
	assert.Equal(t, path, m.GetPath())
	assert.Equal(t, bAddr, m.GetBaseAddress())
}
