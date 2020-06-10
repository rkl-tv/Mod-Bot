package request_handler

import (
	game2 "ModBot/game"
	error2 "ModBot/irc_bot/request_handler/error"
	"ModBot/twitch/irc"
)

type helpRequestHandler struct {
	ircClient     irc.Client
	gameConnector game2.Connector
}

func NewHelpRequestHandler(ircClient irc.Client, gameConnector game2.Connector) RequestHandler {
	return &helpRequestHandler{
		ircClient:     ircClient,
		gameConnector: gameConnector,
	}
}

func (h *helpRequestHandler) Handle(args []string) error {
	game := h.gameConnector.GetGame()
	if game == nil {
		return error2.NewNoGameAttachedError()
	}

	return h.ircClient.Say(game.GetUsage())
}
