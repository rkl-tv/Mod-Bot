package irc_bot

import (
	game2 "ModBot/game"
	"ModBot/irc_bot/input"
	rh "ModBot/irc_bot/request_handler"
	error2 "ModBot/irc_bot/request_handler/error"
	"ModBot/twitch/irc"
	"log"
	"regexp"
	"strings"
)

type DefaultIrcBot struct {
	client          irc.Client
	gameConnector   game2.Connector
	requestHandlers map[string]rh.RequestHandler
}

func NewDefaultIrcBot(client irc.Client, gameConnector game2.Connector) IrcBot {
	b := &DefaultIrcBot{
		client:        client,
		gameConnector: gameConnector,
		requestHandlers: map[string]rh.RequestHandler{
			"help": rh.NewHelpRequestHandler(client, gameConnector),
		},
	}

	return b
}

func (b *DefaultIrcBot) Start() error {
	msgCh, err := b.client.Connect()
	if err != nil {
		return err
	}

	for msg := range msgCh {
		go b.process(msg)
	}

	return nil
}

func (b *DefaultIrcBot) Stop() error {
	return b.client.Disconnect()
}

func (b *DefaultIrcBot) Attach(game game2.Game) {
	b.gameConnector.Attach(game)
}

func (b *DefaultIrcBot) handle(r *input.Request) error {
	expr := regexp.MustCompile("^\\$(([a-z]+)(\\s+([a-zA-Z0-9\\.\\s]+))?)$")
	if !expr.MatchString(r.GetRaw()) {
		u := r.GetIrcUser()

		log.Printf("[M] %s (%s): %s\n", u.DisplayName, u.ID, r.GetRaw())
		return nil
	}

	matches := expr.FindStringSubmatch(r.GetRaw())
	if len(matches) < 3 {
		return nil
	}

	args := strings.Fields(matches[1])
	log.Printf("[R] %s (%s): %s\n", r.GetIrcUser().DisplayName, r.GetIrcUser().ID, args)

	handler, ok := b.requestHandlers[args[0]]
	if !ok {
		return b.forwardToGame(r.GetIrcUser(), args)
	}

	log.Printf("[I] found native handler for \"%v\" request\n", args)
	return handler.Handle(args[1:])
}

func (b *DefaultIrcBot) forwardToGame(ircUser irc.User, args []string) error {
	log.Printf("[I] forward \"%v\" request to attached game\n", args)

	game := b.gameConnector.GetGame()
	if game == nil {
		return error2.NewNoGameAttachedError()
	}

	res, err := game.ProcessRequest(args)
	if err != nil {
		_ = b.client.SayTo(ircUser, err.Error())
		return err
	}

	return b.client.SayTo(ircUser, res.GetMessage())
}

func (b *DefaultIrcBot) process(m irc.PrivateIrcMessage) {
	r := input.NewRequest(m.User, m.Content)

	if err := b.handle(r); err != nil {
		log.Printf("[E] %s\n", err)
	}
}
