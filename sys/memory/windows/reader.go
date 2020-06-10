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

type reader struct {
	pid     process.DWORD
	address uintptr
}

func newReader(pid process.DWORD, address uintptr) io.Reader {
	return &reader{
		pid:     pid,
		address: address,
	}
}

func (r *reader) Read(p []byte) (n int, err error) {
	cLen := C.size_t(len(p))
	cAddress := C.LPCVOID(r.address)

	cProcessHandle := C.OpenProcess(C.PROCESS_ALL_ACCESS, C.FALSE, C.DWORD(r.pid))
	if cProcessHandle == nil {
		return 0, r.newInternalError("couldn't access process")
	}
	defer C.CloseHandle(cProcessHandle)

	var readBytes C.size_t
	if 0 == C.ReadProcessMemory(cProcessHandle, cAddress, (C.LPVOID)(unsafe.Pointer(&p[0])), cLen, &readBytes) {
		return 0, r.newInternalError("couldn't access process")
	}

	return int(readBytes), err
}

func (r *reader) newInternalError(msg string) error {
	code := C.GetLastError()

	return errors.New(fmt.Sprintf("%s (code: %d)", msg, int(code)))
}
