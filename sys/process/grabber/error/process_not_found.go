package error

import "fmt"

type ProcessNotFoundError struct {
	processName string
}

func NewProcessNotFoundError(processName string) error {
	return &ProcessNotFoundError{
		processName: processName,
	}
}

func (e *ProcessNotFoundError) Error() string {
	return fmt.Sprintf("process \"%s\" not found", e.processName)
}
