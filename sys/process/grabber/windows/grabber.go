package windows

// #cgo LDFLAGS: -lpsapi
// #include "find_process_id.h"
// #include "get_all_modules.h"
import "C"

import (
	error2 "ModBot/sys/error"
	"ModBot/sys/process"
	"ModBot/sys/process/grabber"
	"errors"
	"reflect"
	_ "runtime/cgo"
	"unsafe"
)

type Grabber struct {
}

func NewGrabber() grabber.Grabber {
	return &Grabber{}
}

func (g *Grabber) Grab(processName string) (*process.Process, error) {
	cProcessName := C.CString(processName)
	defer C.free(unsafe.Pointer(cProcessName))

	cProcessId := C.findProcessId(cProcessName)
	processId := process.DWORD(cProcessId)

	if process.DWORD(0) == processId {
		return nil, error2.NewProcessNotFoundError(processName)
	}

	modules, err := g.getModulesFor(processId)
	if err != nil {
		return nil, err
	}

	p := process.NewProcess(processId, processName, modules)

	return p, nil
}

func (g *Grabber) getModulesFor(pid process.DWORD) (process.ModuleList, error) {
	var modCount process.DWORD

	cStruct := C.getAllModules(C.DWORD(pid), (*C.DWORD)(unsafe.Pointer(&modCount)))
	if cStruct == nil {
		return nil, errors.New("couldn't load process modules")
	}

	// found here: https://gist.github.com/nasitra/98bb59421be49a518c4a
	var cModList []*C.module_t

	cModListHeader := (*reflect.SliceHeader)(unsafe.Pointer(&cModList))
	cModListHeader.Cap = int(modCount)
	cModListHeader.Len = int(modCount)
	cModListHeader.Data = uintptr(unsafe.Pointer(cStruct))

	resultList := process.ModuleList{}

	for i := process.DWORD(0); i < modCount; i++ {
		cMod := cModList[i]

		m := process.NewModule(C.GoString(cMod.fileName), uintptr(cMod.baseAddr))

		resultList = append(resultList, m)
	}

	return resultList, nil
}
