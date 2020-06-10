package error

import "fmt"

type ModuleBaseAddressNotFoundError struct {
	processId   uint
	processName string
}

func NewModuleBaseAddressNotFoundError(processId uint, processName string) error {
	return &ModuleBaseAddressNotFoundError{
		processId:   processId,
		processName: processName,
	}
}

func (e *ModuleBaseAddressNotFoundError) Error() string {
	return fmt.Sprintf("module base address not found for pid \"%d (%s)\"", e.processId, e.processName)
}
