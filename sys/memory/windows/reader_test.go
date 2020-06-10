package windows

import (
	"ModBot/sys/process"
	"ModBot/sys/process/grabber/windows"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewReader(t *testing.T) {
	pid := process.DWORD(123)
	addr := uintptr(666777)

	r := newReader(pid, addr).(*reader)
	assert.Equal(t, pid, r.pid)
	assert.Equal(t, addr, r.address)
}

func TestReader_Read(t *testing.T) {
	grabber := windows.NewGrabber()

	p, err := grabber.Grab("steam.exe")
	assert.Nil(t, err)

	baseModule := p.GetModules()[0]
	assert.Contains(t, baseModule.GetPath(), "steam.exe")

	r := newReader(p.GetId(), baseModule.GetBaseAddress()) // should read 'MZ'

	buf := make([]byte, 2)
	readBytes, err := r.Read(buf)
	assert.Nil(t, err)
	assert.Equal(t, 2, readBytes)
	assert.Equal(t, []byte{'M', 'Z'}, buf)
}
