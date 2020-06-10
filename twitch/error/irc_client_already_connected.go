package error

import "fmt"

type IrcClientAlreadyConnectedError struct {
}

func NewIrcClientAlreadyConnectedError() error {
	return &IrcClientAlreadyConnectedError{}
}

func (e *IrcClientAlreadyConnectedError) Error() string {
	return fmt.Sprintf("already connected")
}
