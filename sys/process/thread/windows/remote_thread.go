package windows

// #include <windows.h>
import "C"
import (
	"ModBot/sys/process"
	"ModBot/sys/process/thread"
	"errors"
	"fmt"
	"unsafe"
)

type remoteThread struct {
	pid          process.DWORD
	entryAddress uintptr
	argsAddress  *uintptr
}

func newRemoteThread(pid process.DWORD, entryAddress uintptr, argsAddress *uintptr) thread.Remote {
	return &remoteThread{
		pid:          pid,
		entryAddress: entryAddress,
		argsAddress:  argsAddress,
	}
}

func (t *remoteThread) Run() error {
	cFuncAddr := *((*C.LPTHREAD_START_ROUTINE)(unsafe.Pointer(&t.entryAddress)))

	cProcessHandle := C.OpenProcess(C.PROCESS_ALL_ACCESS|C.PROCESS_QUERY_INFORMATION, C.FALSE, C.DWORD(t.pid))
	if cProcessHandle == nil {
		return t.newInternalError("couldn't access process")
	}
	defer C.CloseHandle(cProcessHandle)

	var lpThreadAttributes C.LPSECURITY_ATTRIBUTES
	var lpParameter C.LPVOID
	var dwCreationFlags C.DWORD
	var lpThreadId C.LPDWORD

	rt := C.CreateRemoteThread(cProcessHandle, lpThreadAttributes, 0, cFuncAddr, lpParameter, dwCreationFlags, lpThreadId)
	if rt == nil {
		return t.newInternalError("couldn't exec server module function")
	}
	defer C.CloseHandle(rt)

	return nil
}

func (t *remoteThread) newInternalError(msg string) error {
	code := C.GetLastError()

	return errors.New(fmt.Sprintf("%s (code: %d)", msg, int(code)))
}
