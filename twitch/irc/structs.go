package irc

type User struct {
	ID          string
	Name        string
	DisplayName string
}

type PrivateIrcMessage struct {
	User    User
	Content string
}
