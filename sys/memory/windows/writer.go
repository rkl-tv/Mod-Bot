package windows

// #include <windows.h>
import "C"
import (
	"ModBot/sys/process"
	"errors"
	"fmt"
	"io"
	"unsafe"
)

type writer struct {
	pid     process.DWORD
	address uintptr
}

func newWriter(pid process.DWORD, address uintptr) io.Writer {
	return &writer{
		pid:     pid,
		address: address,
	}
}

// implements io.Writer
func (w *writer) Write(p []byte) (n int, err error) {
	cLen := C.size_t(len(p))
	cAddress := C.LPVOID(w.address)

	cProcessHandle := C.OpenProcess(C.PROCESS_ALL_ACCESS, C.FALSE, C.DWORD(w.pid))
	if cProcessHandle == nil {
		return 0, w.newInternalError("couldn't access process")
	}
	defer C.CloseHandle(cProcessHandle)

	var cOldProtection C.DWORD
	if 0 == C.VirtualProtectEx(cProcessHandle, cAddress, cLen, C.PAGE_READWRITE, &cOldProtection) {
		return 0, w.newInternalError("couldn't access process")
	}

	var writtenBytes C.size_t
	if 0 == C.WriteProcessMemory(cProcessHandle, cAddress, (C.LPCVOID)(unsafe.Pointer(&p[0])), cLen, &writtenBytes) {
		return 0, w.newInternalError("couldn't access process")
	}

	var devNull C.DWORD
	if 0 == C.VirtualProtectEx(cProcessHandle, cAddress, cLen, cOldProtection, &devNull) {
		return 0, w.newInternalError("couldn't access process")
	}

	return int(writtenBytes), nil
}

func (w *writer) newInternalError(msg string) error {
	code := C.GetLastError()

	return errors.New(fmt.Sprintf("%s (code: %d)", msg, int(code)))
}
