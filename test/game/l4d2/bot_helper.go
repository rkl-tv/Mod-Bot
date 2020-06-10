package l4d2

import (
	game2 "ModBot/game"
	"ModBot/game/l4d2"
	"ModBot/irc_bot"
	"ModBot/sys/memory"
	"ModBot/sys/process"
	grabber2 "ModBot/sys/process/grabber"
	"ModBot/sys/process/thread"
	"ModBot/twitch/irc"
	"bytes"
	"fmt"
	"io"
)

type BotHelper struct {
	lastMessages bytes.Buffer
	msgChannel   chan irc.PrivateIrcMessage
}

func NewBotHelper() *BotHelper {
	return &BotHelper{}
}

func (b *BotHelper) GetMessages() string {
	return b.lastMessages.String()
}

func (b *BotHelper) SendPrivateIrcMessage(id, name, displayName, content string) {
	msg := irc.PrivateIrcMessage{
		User: irc.User{
			ID:          id,
			Name:        name,
			DisplayName: displayName,
		},
		Content: content,
	}

	b.msgChannel <- msg
}

func (b *BotHelper) CreateBot() irc_bot.IrcBot {
	b.msgChannel = make(chan irc.PrivateIrcMessage)

	bot := irc_bot.NewDefaultIrcBot(
		b.prepareIrcClient(),
		game2.NewDefaultConnector(),
	)

	bot.Attach(b.prepareGame())

	return bot
}

func (b *BotHelper) prepareIrcClient() irc.Client {
	client := &irc.MockedClient{}

	client.ConnectFunc = func() (chan irc.PrivateIrcMessage, error) {
		return b.msgChannel, nil
	}

	client.DisconnectFunc = func() error {
		close(b.msgChannel)
		return nil
	}

	client.SayFunc = func(message string) error {
		msg := irc.PrivateIrcMessage{
			User: irc.User{
				ID:          "666",
				Name:        "rkl85",
				DisplayName: "RKL85",
			},
			Content: message,
		}

		b.lastMessages.WriteString(msg.Content)
		b.msgChannel <- msg

		return nil
	}

	client.SayToFunc = func(user irc.User, message string) error {
		msg := irc.PrivateIrcMessage{
			User: irc.User{
				ID:          "666",
				Name:        "rkl85",
				DisplayName: "RKL85",
			},
			Content: fmt.Sprintf("@%s %s", user.DisplayName, message),
		}

		b.lastMessages.WriteString(msg.Content)
		b.msgChannel <- msg

		return nil
	}

	return client
}

func (b *BotHelper) prepareGame() game2.Game {
	grabber := &grabber2.MockedGrabber{}
	memWriterFactory := &memory.MockedWriterFactory{}
	remoteThreadFactory := &thread.MockedRemoteFactory{}

	memWriterFactory.NewFunc = func(process interface{}, address uintptr) io.Writer {
		buf := bytes.Buffer{}
		return &buf
	}

	remoteThreadFactory.NewFunc = func(process interface{}, entryAddress uintptr, argsAddress *uintptr) thread.Remote {
		return &thread.MockedRemote{}
	}

	grabber.GrabFunc = func(processName string) (*process.Process, error) {
		mod := process.NewModule("/foo/bar/server.dll", uintptr(0x123456))
		prc := process.NewProcess(process.DWORD(666), "left4dead2.exe", process.ModuleList{mod})

		return prc, nil
	}

	config := l4d2.NewConfig()
	config.SetBoostSeconds(1)

	return l4d2.NewL4D2(grabber, memWriterFactory, remoteThreadFactory, config)
}
