package irc

import (
	error2 "ModBot/twitch/irc/error"
	"github.com/gempir/go-twitch-irc/v2"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDefaultIrcClient(t *testing.T) {
	username := "foobar"
	oauth := "123456"
	channel := "ch"

	b := NewDefaultClient(username, oauth, channel).(*DefaultClient)
	assert.Equal(t, username, b.username)
	assert.Equal(t, oauth, b.oauthCredentials)
	assert.Equal(t, channel, b.channel)
}

func TestDefaultClient_Connect__irc_client_already_connected_error(t *testing.T) {
	b := NewDefaultClient("", "", "").(*DefaultClient)
	b.client = twitch.NewClient("foo", "bar")

	_, err := b.Connect()
	assert.IsType(t, &error2.IrcClientAlreadyConnectedError{}, err)
}

func TestDefaultClient_Connect__success(t *testing.T) {
	c := NewDefaultClient("", "", "").(*DefaultClient)
	c.thirdPartyConnectFunc = func() error {
		return nil
	}

	inputCh, err := c.Connect()
	assert.Nil(t, err)
	assert.NotNil(t, inputCh)
}

func TestDefaultClient_Connect__error(t *testing.T) {
	//defer func() {
	//	if r := recover(); r == nil {
	//		t.Errorf("The code did not panic")
	//	}
	//}()
	//
	//c := NewDefaultClient("u", "o", "c").(*DefaultClient)
	//c.thirdPartyConnectFunc = func() error {
	//	return errors.New("test error")
	//}
	//
	//_, err := c.Connect()
	//assert.Nil(t, err)
}

func TestDefaultClient_Disconnect__irc_client_not_connected_error(t *testing.T) {
	c := NewDefaultClient("", "", "")

	err := c.Disconnect()
	assert.IsType(t, &error2.IrcClientNotConnectedError{}, err)
}

func TestDefaultClient_Disconnect__success(t *testing.T) {
	c := NewDefaultClient("", "", "").(*DefaultClient)
	c.thirdPartyConnectFunc = func() error {
		return nil
	}
	c.thirdPartyDisconnectFunc = func() error {
		return nil
	}

	var err error

	_, err = c.Connect()
	assert.Nil(t, err)

	c.client = twitch.NewClient("foo", "bar")
	c.messagesChannel = make(chan PrivateIrcMessage)

	err = c.Disconnect()
	assert.Nil(t, err)
	assert.Nil(t, c.client)
	assert.Nil(t, c.messagesChannel)
}

func TestDefaultIrcClient_SayTo(t *testing.T) {
	c := NewDefaultClient("", "", "test-channel").(*DefaultClient)

	var captChannel string
	var captText string

	c.thirdPartySayFunc = func(channel, text string) {
		captChannel = channel
		captText = text
	}

	err := c.SayTo(User{DisplayName: "Test-User"}, "Test message")
	assert.Nil(t, err)
	assert.Equal(t, "test-channel", captChannel)
	assert.Equal(t, "@Test-User Test message", captText)
}

func TestDefaultClient_Say(t *testing.T) {
	c := NewDefaultClient("", "", "test-channel").(*DefaultClient)

	var captChannel string
	var captText string

	c.thirdPartySayFunc = func(channel, text string) {
		captChannel = channel
		captText = text
	}

	err := c.Say("Test message")
	assert.Nil(t, err)
	assert.Equal(t, "test-channel", captChannel)
	assert.Equal(t, "Test message", captText)
}
