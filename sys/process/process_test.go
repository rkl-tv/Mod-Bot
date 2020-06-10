package process

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewProcess(t *testing.T) {
	id := DWORD(123)
	name := "foobar.exe"
	modules := ModuleList{NewModule("", uintptr(0xBBEEFF))}

	p := NewProcess(id, name, modules)
	assert.NotNil(t, p)
	assert.Equal(t, id, p.GetId())
	assert.Equal(t, name, p.GetName())
	assert.Equal(t, modules, p.GetModules())
}
