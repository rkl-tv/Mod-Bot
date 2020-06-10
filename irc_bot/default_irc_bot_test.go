package irc_bot

import (
	game2 "ModBot/game"
	"ModBot/irc_bot/input"
	error2 "ModBot/irc_bot/request_handler/error"
	"ModBot/twitch/irc"
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestNewDefaultIrcBot(t *testing.T) {
	client := &irc.MockedClient{}
	connector := &game2.DefaultConnector{}

	bot := NewDefaultIrcBot(client, connector).(*DefaultIrcBot)
	assert.Equal(t, client, bot.client)
	assert.Equal(t, connector, bot.gameConnector)
}

func TestDefaultIrcBot_Start(t *testing.T) {
	mockedIrcClient := &irc.MockedClient{}

	msgCh := make(chan irc.PrivateIrcMessage)
	mockedIrcClient.ConnectFunc = func() (chan irc.PrivateIrcMessage, error) {
		return msgCh, nil
	}

	b := NewDefaultIrcBot(mockedIrcClient, &game2.DefaultConnector{}).(*DefaultIrcBot)

	go func() {
		err := b.Start()
		assert.Nil(t, err)
	}()

	time.Sleep(1 * time.Second)
	close(msgCh)
}

func TestDefaultIrcBot_Stop(t *testing.T) {
	mockedIrcClient := &irc.MockedClient{}

	msgCh := make(chan irc.PrivateIrcMessage)

	mockedIrcClient.ConnectFunc = func() (chan irc.PrivateIrcMessage, error) {
		return msgCh, nil
	}

	mockedIrcClient.DisconnectFunc = func() error {
		close(msgCh)
		return nil
	}

	b := NewDefaultIrcBot(mockedIrcClient, &game2.DefaultConnector{}).(*DefaultIrcBot)

	hasStopped := false
	go func() {
		_ = b.Start()
		hasStopped = true
	}()

	err := b.Stop()
	assert.Nil(t, err)

	time.Sleep(1 * time.Second)
	assert.True(t, hasStopped)
}

func TestDefaultIrcBot_Attach(t *testing.T) {
	b := NewDefaultIrcBot(&irc.MockedClient{}, &game2.DefaultConnector{}).(*DefaultIrcBot)
	assert.Nil(t, b.gameConnector.GetGame())

	game := &game2.MockedGame{}
	b.Attach(game)
	assert.Equal(t, game, b.gameConnector.GetGame())
}

func TestNewDefaultIrcBot_handle(t *testing.T) {
	client := &irc.MockedClient{}
	connector := game2.NewDefaultConnector()
	b := NewDefaultIrcBot(client, connector).(*DefaultIrcBot)

	// ignore pattern mismatch
	{
		var lBuf bytes.Buffer
		log.SetOutput(&lBuf)

		err := b.handle(input.NewRequest(irc.User{ID: "123", Name: "rkl85", DisplayName: "RKL85"}, "some text with ö"))
		assert.Nil(t, err)
		assert.Contains(t, lBuf.String(), "[M] RKL85 (123): some text with ö\n")
	}

	// forward to game, when no native handler was found
	{
		var lBuf bytes.Buffer
		log.SetOutput(&lBuf)

		_ = b.handle(input.NewRequest(irc.User{ID: "bar", Name: "", DisplayName: "foo"}, "$game command"))
		assert.Contains(t, lBuf.String(), "[R] foo (bar): [game command]\n")
		assert.Contains(t, lBuf.String(), "[I] forward \"[game command]\" request to attached game\n")
	}

	// native handler was found
	{
		var lBuf bytes.Buffer
		log.SetOutput(&lBuf)

		_ = b.handle(input.NewRequest(irc.User{ID: "foo", Name: "", DisplayName: "bar"}, "$valid request with some details"))
		assert.Contains(t, lBuf.String(), "[R] bar (foo): [valid request with some details]\n")
		assert.Contains(t, lBuf.String(), "[I] forward \"[valid request with some details]\" request to attached game\n")
	}
}

func TestDefaultIrcBot_forwardToGame(t *testing.T) {
	client := &irc.MockedClient{}
	connector := game2.NewDefaultConnector()
	b := NewDefaultIrcBot(client, connector).(*DefaultIrcBot)

	// NoGameAttachedError
	{
		err := b.forwardToGame(irc.User{}, []string{})
		assert.IsType(t, &error2.NoGameAttachedError{}, err)
	}

	// call game processor with error result
	{
		mg := &game2.MockedGame{}
		mg.ProcessRequestFunc = func(args []string) (*game2.Response, error) {
			return game2.NewResponse(""), errors.New("test error")
		}
		b.gameConnector.Attach(mg)

		err := b.forwardToGame(irc.User{}, []string{})
		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	}

	// call game processor with non error result
	{
		captMessage := ""
		client.SayToFunc = func(user irc.User, message string) error {
			captMessage = fmt.Sprintf("@%s %s", user.DisplayName, message)
			return nil
		}

		mg := &game2.MockedGame{}
		mg.ProcessRequestFunc = func(args []string) (*game2.Response, error) {
			return game2.NewResponse("hello my friend"), nil
		}
		b.gameConnector.Attach(mg)

		err := b.forwardToGame(irc.User{DisplayName: "Test"}, []string{})
		assert.Nil(t, err)
		assert.Equal(t, "@Test hello my friend", captMessage)
	}
}

func TestDefaultIrcBot_process(t *testing.T) {
	client := &irc.MockedClient{}
	connector := game2.NewDefaultConnector()
	b := NewDefaultIrcBot(client, connector).(*DefaultIrcBot)

	// handle request error
	{
		var lBuf bytes.Buffer
		log.SetOutput(&lBuf)

		b.process(irc.PrivateIrcMessage{User: irc.User{ID: "123", Name: "foo", DisplayName: "FOO"}, Content: "$foobar"})
		assert.Contains(t, lBuf.String(), "[E] no game attached\n")
	}
}
