//+build wireinject

package main

import (
	"ModBot/app"
	"ModBot/config"
	"ModBot/game"
	"ModBot/game/l4d2"
	"ModBot/irc_bot"
	"ModBot/sys/memory"
	windows2 "ModBot/sys/memory/windows"
	"ModBot/sys/process/grabber"
	"ModBot/sys/process/grabber/windows"
	"ModBot/sys/process/thread"
	windows3 "ModBot/sys/process/thread/windows"
	"ModBot/twitch/irc"
	"github.com/google/wire"
)

func InitTwitchIrcClient(cfg config.Config) irc.Client {
	return irc.NewDefaultClient(
		cfg.GetTwitchIrcUsername(),
		cfg.GetTwitchIrcAuthentication(),
		cfg.GetTwitchIrcChannel(),
	)
}

func InitIrcBot(ircClient irc.Client, connector game.Connector) irc_bot.IrcBot {
	return irc_bot.NewDefaultIrcBot(ircClient, connector)
}

func InitL4D2Game(
	processGrabber grabber.Grabber,
	memWriterFactory memory.WriterFactory,
	remoteThreadFactory thread.RemoteFactory,
	cfg config.Config,
) game.Game {
	l4d2Config := l4d2.NewConfig()
	l4d2Config.SetBoostSeconds(cfg.L4D2GetBoostSeconds())

	return l4d2.NewL4D2(processGrabber, memWriterFactory, remoteThreadFactory, l4d2Config)
}

func ProvideL4D2() game.Game {
	wire.Build(InitL4D2Game, ProvideProcessGrabber, ProvideMemWriterFactory, ProvideRemoteThreadFactory, ProvideConfig)
	return &game.MockedGame{}
}

func ProvideRemoteThreadFactory() thread.RemoteFactory {
	wire.Build(windows3.NewRemoteThreadFactory)
	return windows3.NewRemoteThreadFactory()
}

func ProvideMemWriterFactory() memory.WriterFactory {
	wire.Build(windows2.NewMemWriterFactory)
	return windows2.NewMemWriterFactory()
}

func ProvideProcessGrabber() grabber.Grabber {
	wire.Build(windows.NewGrabber)
	return windows.NewGrabber()
}

func ProvideConfig() config.Config {
	wire.Build(config.NewIniConfig)
	return &config.IniConfig{}
}

func ProvideTwitchIrcClient() irc.Client {
	wire.Build(InitTwitchIrcClient, ProvideConfig)
	return &irc.DefaultClient{}
}

func ProvideGameConnector() game.Connector {
	wire.Build(game.NewDefaultConnector)
	return &game.DefaultConnector{}
}

func ProvideIrcBot() irc_bot.IrcBot {
	wire.Build(InitIrcBot, ProvideTwitchIrcClient, ProvideGameConnector)
	return &irc_bot.DefaultIrcBot{}
}

func ProvideApp() app.App {
	wire.Build(app.NewDefaultApp, ProvideIrcBot, ProvideL4D2)
	return &app.DefaultApp{}
}
