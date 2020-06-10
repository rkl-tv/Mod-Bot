package irc_bot

import "ModBot/game"

// IrcBot Twitch irc bot
type IrcBot interface {
	Start() error
	Stop() error
	Attach(game game.Game)
}
