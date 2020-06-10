package irc

import (
	error2 "ModBot/twitch/irc/error"
	"fmt"
	"github.com/gempir/go-twitch-irc/v2"
)

type DefaultClient struct {
	username                 string
	oauthCredentials         string
	channel                  string
	client                   *twitch.Client
	messagesChannel          chan PrivateIrcMessage
	thirdPartyConnectFunc    func() error
	thirdPartyDisconnectFunc func() error
	thirdPartySayFunc        func(channel, text string)
}

func NewDefaultClient(username, oauthCredentials, channel string) Client {
	c := &DefaultClient{
		username:         username,
		oauthCredentials: oauthCredentials,
		channel:          channel,
	}

	c.thirdPartyConnectFunc = func() error {
		return c.client.Connect()
	}

	c.thirdPartyDisconnectFunc = func() error {
		return c.client.Disconnect()
	}

	c.thirdPartySayFunc = func(channel, text string) {
		c.client.Say(channel, text)
	}

	return c
}

func (c *DefaultClient) Say(message string) error {
	c.thirdPartySayFunc(c.channel, message)
	return nil
}

func (c *DefaultClient) SayTo(user User, message string) error {
	c.thirdPartySayFunc(c.channel, fmt.Sprintf("@%s %s", user.DisplayName, message))
	return nil
}

func (c *DefaultClient) Connect() (chan PrivateIrcMessage, error) {
	if nil != c.client {
		return nil, error2.NewIrcClientAlreadyConnectedError()
	}

	c.messagesChannel = make(chan PrivateIrcMessage)

	c.client = twitch.NewClient(c.username, c.oauthCredentials)
	c.client.Join(c.channel)

	c.client.OnPrivateMessage(func(message twitch.PrivateMessage) {
		prvIrcMsg := PrivateIrcMessage{
			User: User{
				ID:          message.User.ID,
				Name:        message.User.Name,
				DisplayName: message.User.DisplayName,
			},
			Content: message.Message,
		}
		c.messagesChannel <- prvIrcMsg
	})

	go func() {
		defer func() {
			c.client = nil
			if nil != c.messagesChannel {
				close(c.messagesChannel)
			}
		}()

		if err := c.thirdPartyConnectFunc(); err != nil {
			panic(err)
		}
	}()

	return c.messagesChannel, nil
}

func (c *DefaultClient) Disconnect() error {
	if nil == c.client {
		return error2.NewIrcClientNotConnectedError()
	}

	close(c.messagesChannel)

	c.client = nil
	c.messagesChannel = nil

	return c.thirdPartyDisconnectFunc()
}
