package error

type NoGameAttachedError struct {
}

func NewNoGameAttachedError() error {
	return &NoGameAttachedError{}
}

func (e *NoGameAttachedError) Error() string {
	return "no game attached"
}
