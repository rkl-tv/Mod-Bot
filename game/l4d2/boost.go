package l4d2

// #include <windows.h>
import "C"
import (
	error2 "ModBot/game/l4d2/error"
	"ModBot/sys/memory"
	process2 "ModBot/sys/process"
	"ModBot/sys/process/grabber"
	"ModBot/sys/process/thread"
	"fmt"
	"log"
	"strings"
	"time"
)

type boostCommand struct {
	processName         string
	processGrabber      grabber.Grabber
	boostSeconds        uint
	memWriterFactory    memory.WriterFactory
	remoteThreadFactory thread.RemoteFactory

	enableFunc               func(pid process2.DWORD, baseAddr uintptr) error
	disableFunc              func(pid process2.DWORD, baseAddr uintptr) error
	patchServerModuleFunc    func(value byte, pid process2.DWORD, baseAddr uintptr) error
	execServerModuleCodeFunc func(pid process2.DWORD, baseAddress uintptr) error

	active bool
}

func newBoostCommand(
	processName string,
	seconds uint,
	processGrabber grabber.Grabber,
	memWriterFactory memory.WriterFactory,
	remoteThreadFactory thread.RemoteFactory,
) *boostCommand {
	c := &boostCommand{
		processName:         processName,
		processGrabber:      processGrabber,
		boostSeconds:        seconds,
		memWriterFactory:    memWriterFactory,
		remoteThreadFactory: remoteThreadFactory,
		active:              false,
	}

	c.enableFunc = c.enable
	c.disableFunc = c.disable
	c.patchServerModuleFunc = c.patchServerModule
	c.execServerModuleCodeFunc = c.execServerModuleCode

	return c
}

func (c *boostCommand) boost() error {
	log.Printf(fmt.Sprintf("[I] Initiate team boost for %d seconds\n", c.boostSeconds))

	if c.active {
		return error2.NewBoostIsActiveError()
	}

	process, err := c.processGrabber.Grab(c.processName)
	if err != nil {
		return err
	}

	var mod *process2.Module

	for _, m := range process.GetModules() {
		if strings.Contains(m.GetPath(), "server.dll") {
			mod = m
			break
		}
	}

	if nil == mod {
		return error2.NewServerModuleNotFoundError()
	}

	// async invoke for timeout
	var enableErr error

	go func(p process2.DWORD, ba uintptr) {
		if enableErr = c.enableFunc(p, ba); enableErr != nil {
			log.Printf("[E] boost error: %s\n", enableErr)
		}
	}(process.GetId(), mod.GetBaseAddress())

	// start timer to disable boost
	time.Sleep(time.Duration(c.boostSeconds) * time.Second)

	if enableErr != nil {
		return enableErr
	}

	return c.disableFunc(process.GetId(), mod.GetBaseAddress())
}

func (c *boostCommand) enable(pid process2.DWORD, baseAddr uintptr) error {
	// 0x31 = ASCII('1')
	if err := c.patchServerModuleFunc(0x31, pid, baseAddr); err != nil {
		return err
	}

	if err := c.execServerModuleCodeFunc(pid, baseAddr); err != nil {
		return err
	}

	c.active = true

	return nil
}

func (c *boostCommand) disable(pid process2.DWORD, baseAddr uintptr) error {
	defer func() { c.active = false }()

	// 0x30 = ASCII('0')
	if err := c.patchServerModuleFunc(0x30, pid, baseAddr); err != nil {
		return err
	}

	if err := c.execServerModuleCodeFunc(pid, baseAddr); err != nil {
		return err
	}

	return nil
}

func (c *boostCommand) patchServerModule(value byte, pid process2.DWORD, baseAddr uintptr) error {
	patchAddress := baseAddr + 0x5AA0A4

	writer := c.memWriterFactory.New(pid, patchAddress)

	_, err := writer.Write([]byte{value})
	if err != nil {
		return err
	}

	return nil
}

func (c *boostCommand) execServerModuleCode(pid process2.DWORD, baseAddress uintptr) error {
	funcAddress := baseAddress + 0x5817B0

	remoteThread := c.remoteThreadFactory.New(pid, funcAddress, nil)

	return remoteThread.Run()
}
