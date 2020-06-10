package irc

type MockedClient struct {
	ConnectFunc    func() (chan PrivateIrcMessage, error)
	DisconnectFunc func() error
	SayToFunc      func(user User, message string) error
	SayFunc        func(message string) error
}

func (c *MockedClient) Connect() (chan PrivateIrcMessage, error) {
	if nil != c.ConnectFunc {
		return c.ConnectFunc()
	}

	return nil, nil
}

func (c *MockedClient) Disconnect() error {
	if nil != c.DisconnectFunc {
		return c.DisconnectFunc()
	}

	return nil
}

func (c *MockedClient) SayTo(user User, message string) error {
	if nil != c.SayToFunc {
		return c.SayToFunc(user, message)
	}

	return nil
}

func (c *MockedClient) Say(message string) error {
	if nil != c.SayFunc {
		return c.SayFunc(message)
	}

	return nil
}
