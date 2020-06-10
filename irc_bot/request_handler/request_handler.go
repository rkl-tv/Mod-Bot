package request_handler

type RequestHandler interface {
	Handle(args []string) error
}
