package app

import (
	"ModBot/game"
	"ModBot/irc_bot"
)

type DefaultApp struct {
	bot  irc_bot.IrcBot
	l4d2 game.Game
}

func NewDefaultApp(bot irc_bot.IrcBot, l4d2 game.Game) App {
	return &DefaultApp{
		bot:  bot,
		l4d2: l4d2,
	}
}

func (a *DefaultApp) Run() error {
	a.bot.Attach(a.l4d2)

	return a.bot.Start()
}
