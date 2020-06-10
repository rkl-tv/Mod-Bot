package error

type BoostIsActiveError struct {
}

func NewBoostIsActiveError() error {
	return &BoostIsActiveError{}
}

func (e *BoostIsActiveError) Error() string {
	return "boost is already active"
}
