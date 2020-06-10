package l4d2

import (
	"ModBot/game"
	"ModBot/sys/memory"
	"ModBot/sys/process/grabber"
	"ModBot/sys/process/thread"
)

const processName = "left4dead2.exe"

type l4d2 struct {
	processName         string
	processGrabber      grabber.Grabber
	boostCommand        *boostCommand
	memWriterFactory    memory.WriterFactory
	remoteThreadFactory thread.RemoteFactory
	config              *Config
}

func NewL4D2(
	processGrabber grabber.Grabber,
	memWriterFactory memory.WriterFactory,
	remoteThreadFactory thread.RemoteFactory,
	config *Config,
) game.Game {
	g := &l4d2{
		processName:         processName,
		processGrabber:      processGrabber,
		memWriterFactory:    memWriterFactory,
		remoteThreadFactory: remoteThreadFactory,
		config:              config,
	}

	g.boostCommand = newBoostCommand(
		g.processName,
		config.GetBoostSeconds(),
		g.processGrabber,
		memWriterFactory,
		remoteThreadFactory,
	)

	return g
}

func (g *l4d2) GetUsage() string {
	return l4d2Usage
}

func (g *l4d2) ProcessRequest(args []string) (*game.Response, error) {
	switch args[0] {
	case "boost":
		return game.NewResponse("your boost has finished"), g.boostCommand.boost()
	default:
		return game.NewResponse(g.GetUsage()), nil
	}
}
