package game

type Response struct {
	message string
}

func NewResponse(message string) *Response {
	return &Response{
		message: message,
	}
}

func (r *Response) GetMessage() string {
	return r.message
}
