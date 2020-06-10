package game

type Game interface {
	GetUsage() string
	ProcessRequest(args []string) (*Response, error)
}
