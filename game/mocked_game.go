package game

type MockedGame struct {
	GetUsageFunc       func() string
	ProcessRequestFunc func(args []string) (*Response, error)
}

func (g *MockedGame) GetUsage() string {
	if g.GetUsageFunc != nil {
		return g.GetUsageFunc()
	}

	return ""
}

func (g *MockedGame) ProcessRequest(args []string) (*Response, error) {
	if g.ProcessRequestFunc != nil {
		return g.ProcessRequestFunc(args)
	}

	return NewResponse(""), nil
}
