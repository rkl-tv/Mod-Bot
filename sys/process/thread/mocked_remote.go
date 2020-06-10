package thread

type MockedRemote struct {
	RunFunc func() error
}

func (m *MockedRemote) Run() error {
	if m.RunFunc != nil {
		return m.RunFunc()
	}

	return nil
}
