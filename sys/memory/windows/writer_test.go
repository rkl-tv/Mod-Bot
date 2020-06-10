package windows

import (
	"ModBot/sys/process"
	"ModBot/sys/process/grabber/windows"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewWriter(t *testing.T) {
	p := process.DWORD(666)
	a := uintptr(123)

	w := newWriter(p, a).(*writer)
	assert.NotNil(t, w)
	assert.Equal(t, p, w.pid)
	assert.Equal(t, a, w.address)
}

func TestWriter_Write(t *testing.T) {
	grabber := windows.NewGrabber()

	p, err := grabber.Grab("steam.exe")
	assert.Nil(t, err)

	baseModule := p.GetModules()[0]
	assert.Contains(t, baseModule.GetPath(), "steam.exe")

	writer := newWriter(p.GetId(), baseModule.GetBaseAddress()+0x1) // Should be the 'Z' in the first 'MZ' string
	b, err := writer.Write([]byte{'M'})                             // should replace 'MZ' to 'MM'
	defer func() { _, _ = writer.Write([]byte{'Z'}) }()             // reset
	assert.Nil(t, err)
	assert.Equal(t, 1, b)

	// verify written data
	reader := newReader(p.GetId(), baseModule.GetBaseAddress()) // should be 'MM' now

	buf := make([]byte, 2)
	rb, err := reader.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, 2, rb)
	assert.Equal(t, []byte{'M', 'M'}, buf)
}
