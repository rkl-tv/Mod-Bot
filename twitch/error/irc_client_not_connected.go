package error

import "fmt"

type IrcClientNotConnectedError struct {
}

func NewIrcClientNotConnectedError() error {
	return &IrcClientNotConnectedError{}
}

func (e *IrcClientNotConnectedError) Error() string {
	return fmt.Sprintf("not connected")
}
