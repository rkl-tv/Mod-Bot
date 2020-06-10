package windows

import (
	"ModBot/sys/process"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoteThread_newRemoteThread(t *testing.T) {
	pid := process.DWORD(666)
	ea := uintptr(123)
	aa := uintptr(321)

	rt := newRemoteThread(pid, ea, &aa).(*remoteThread)
	assert.NotNil(t, rt)
	assert.Equal(t, pid, rt.pid)
	assert.Equal(t, ea, rt.entryAddress)
	assert.Equal(t, &aa, rt.argsAddress)
}

func TestRemoteThread_Run(t *testing.T) {
	// TODO, currently no idea how to test it well, but I can confirm that it works for L4D2 boost command.
}
