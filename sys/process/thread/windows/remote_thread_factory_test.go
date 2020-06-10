package windows

import (
	"ModBot/sys/process"
	"ModBot/sys/process/thread"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoteThreadFactory_New(t *testing.T) {
	f := NewRemoteThreadFactory()
	argsAddr := uintptr(0x321)

	rt := f.New(process.DWORD(666), uintptr(0x123), &argsAddr)
	assert.IsType(t, &remoteThread{}, rt)
	assert.Implements(t, (*thread.Remote)(nil), rt)
}
