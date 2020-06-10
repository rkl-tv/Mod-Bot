package error

type ServerModuleNotFoundError struct {
}

func NewServerModuleNotFoundError() error {
	return &ServerModuleNotFoundError{}
}

func (e *ServerModuleNotFoundError) Error() string {
	return "server module not found"
}
