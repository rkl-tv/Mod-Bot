package windows

import (
	process2 "ModBot/sys/process"
	"ModBot/sys/process/thread"
)

type RemoteThreadFactory struct {
}

func NewRemoteThreadFactory() thread.RemoteFactory {
	return &RemoteThreadFactory{}
}

func (f *RemoteThreadFactory) New(process interface{}, entryAddress uintptr, argsAddress *uintptr) thread.Remote {
	pid, _ := process.(process2.DWORD) // TODO, handle panic

	return newRemoteThread(pid, entryAddress, argsAddress)
}
