package irc

type Client interface {
	Connect() (chan PrivateIrcMessage, error)
	Disconnect() error
	Say(message string) error
	SayTo(user User, message string) error
}
