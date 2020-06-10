package request_handler

import (
	"ModBot/game"
	error2 "ModBot/irc_bot/request_handler/error"
	"ModBot/twitch/irc"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewHelpRequestHandler(t *testing.T) {
	ic := &irc.MockedClient{}
	gc := game.NewDefaultConnector()

	h := NewHelpRequestHandler(ic, gc).(*helpRequestHandler)
	assert.NotNil(t, h)
	assert.Equal(t, ic, h.ircClient)
	assert.Equal(t, gc, h.gameConnector)
}

func TestHelpRequestHandler_Handle(t *testing.T) {
	// NewNoGameAttachedError
	{
		h := NewHelpRequestHandler(&irc.MockedClient{}, game.NewDefaultConnector())

		err := h.Handle([]string{})
		assert.IsType(t, &error2.NoGameAttachedError{}, err)
	}

	// Print game usage to chat
	{
		ic := &irc.MockedClient{}
		gc := game.NewDefaultConnector()
		mg := &game.MockedGame{}

		captUsage := ""
		ic.SayFunc = func(message string) error {
			captUsage = message
			return nil
		}

		mg.GetUsageFunc = func() string {
			return "dummy usage"
		}

		gc.Attach(mg)

		h := NewHelpRequestHandler(ic, gc).(*helpRequestHandler)

		err := h.Handle([]string{})
		assert.Nil(t, err)
		assert.Equal(t, "dummy usage", captUsage)
	}
}
