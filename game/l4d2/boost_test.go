package l4d2

import (
	error2 "ModBot/game/l4d2/error"
	"ModBot/sys/memory"
	"ModBot/sys/memory/windows"
	"ModBot/sys/process"
	grabber2 "ModBot/sys/process/grabber"
	"ModBot/sys/process/thread"
	windows2 "ModBot/sys/process/thread/windows"
	"bytes"
	"errors"
	"github.com/stretchr/testify/assert"
	"io"
	"log"
	"testing"
)

func TestL4d2_newBoostCommand(t *testing.T) {
	pn := "foobar"
	mg := &grabber2.MockedGrabber{}
	mwf := windows.NewMemWriterFactory()
	rtf := windows2.NewRemoteThreadFactory()
	s := uint(1)

	c := newBoostCommand(pn, s, mg, mwf, rtf)
	assert.NotNil(t, c)
	assert.Equal(t, pn, c.processName)
	assert.Equal(t, mg, c.processGrabber)
	assert.Equal(t, s, c.boostSeconds)
	assert.Equal(t, mwf, c.memWriterFactory)
	assert.Equal(t, rtf, c.remoteThreadFactory)
	assert.Equal(t, false, c.active)
}

func TestL4d2_boost(t *testing.T) {
	// grabber error
	{
		mockedGrabber := &grabber2.MockedGrabber{}
		c := newBoostCommand(processName, uint(1), mockedGrabber, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory())
		c.boostSeconds = 1

		mockedGrabber.GrabFunc = func(processName string) (*process.Process, error) {
			return nil, errors.New("test error")
		}

		err := c.boost()
		assert.NotNil(t, err)
		assert.Equal(t, "test error", err.Error())
	}

	// BoostIsActiveError
	{
		mockedGrabber := &grabber2.MockedGrabber{}
		c := newBoostCommand(processName, uint(1), mockedGrabber, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory())
		c.boostSeconds = 1
		c.active = true

		err := c.boost()
		assert.IsType(t, &error2.BoostIsActiveError{}, err)
	}

	// ServerModuleNotFoundError
	{
		mockedGrabber := &grabber2.MockedGrabber{}
		c := newBoostCommand(processName, uint(1), mockedGrabber, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory())
		c.boostSeconds = 1

		mockedGrabber.GrabFunc = func(processName string) (*process.Process, error) {
			return process.NewProcess(1, processName, process.ModuleList{}), nil
		}

		err := c.boost()
		assert.NotNil(t, err)
		assert.IsType(t, &error2.ServerModuleNotFoundError{}, err)
	}

	// enable error
	{
		mockedGrabber := &grabber2.MockedGrabber{}
		c := newBoostCommand(processName, uint(1), mockedGrabber, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory())
		c.boostSeconds = 1

		mockedGrabber.GrabFunc = func(processName string) (*process.Process, error) {
			mockedModule := process.NewModule("/foo/bar/server.dll", 123456)
			return process.NewProcess(1, processName, process.ModuleList{mockedModule}), nil
		}

		c.enableFunc = func(process.DWORD, uintptr) error {
			return errors.New("enable error")
		}

		isDisabled := false
		c.disableFunc = func(process.DWORD, uintptr) error {
			isDisabled = true
			return nil
		}

		var lBuf bytes.Buffer
		log.SetOutput(&lBuf)

		err := c.boost()
		assert.NotNil(t, err)
		assert.Equal(t, "enable error", err.Error())
		assert.Contains(t, lBuf.String(), "[E] boost error: enable error\n")
		assert.False(t, isDisabled)
	}

	// success test
	{
		mockedGrabber := &grabber2.MockedGrabber{}
		c := newBoostCommand(processName, uint(1), mockedGrabber, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory())
		c.boostSeconds = 1

		mockedGrabber.GrabFunc = func(processName string) (*process.Process, error) {
			mockedModule := process.NewModule("/foo/bar/server.dll", 123456)
			return process.NewProcess(1, processName, process.ModuleList{mockedModule}), nil
		}

		c.enableFunc = func(process.DWORD, uintptr) error {
			return nil
		}

		isDisabled := false
		c.disableFunc = func(process.DWORD, uintptr) error {
			isDisabled = true
			return nil
		}

		err := c.boost()
		assert.Nil(t, err)
		assert.True(t, isDisabled)
	}
}

func TestL4d2_boost_enable(t *testing.T) {
	// only set active, if all steps passed
	c := &boostCommand{}
	c.active = false

	c.patchServerModuleFunc = func(value byte, pid process.DWORD, baseAddr uintptr) error { return errors.New("test error") }
	c.execServerModuleCodeFunc = func(pid process.DWORD, baseAddress uintptr) error { return errors.New("test error") }
	_ = c.enable(process.DWORD(123), uintptr(0x123456))
	assert.False(t, c.active)

	c.patchServerModuleFunc = func(value byte, pid process.DWORD, baseAddr uintptr) error { return nil }
	c.execServerModuleCodeFunc = func(pid process.DWORD, baseAddress uintptr) error { return errors.New("test error") }
	_ = c.enable(process.DWORD(123), uintptr(0x123456))
	assert.False(t, c.active)

	c.patchServerModuleFunc = func(value byte, pid process.DWORD, baseAddr uintptr) error { return errors.New("test error") }
	c.execServerModuleCodeFunc = func(pid process.DWORD, baseAddress uintptr) error { return nil }
	_ = c.enable(process.DWORD(123), uintptr(0x123456))
	assert.False(t, c.active)

	c.patchServerModuleFunc = func(value byte, pid process.DWORD, baseAddr uintptr) error { return nil }
	c.execServerModuleCodeFunc = func(pid process.DWORD, baseAddress uintptr) error { return nil }
	_ = c.enable(process.DWORD(123), uintptr(0x123456))
	assert.True(t, c.active)
}

func TestL4d2_boost_disable(t *testing.T) {
	c := &boostCommand{}
	c.patchServerModuleFunc = func(value byte, pid process.DWORD, baseAddr uintptr) error { return nil }
	c.execServerModuleCodeFunc = func(pid process.DWORD, baseAddress uintptr) error { return errors.New("test error") }
	c.active = true

	_ = c.disable(process.DWORD(123), uintptr(0x123456))
	assert.False(t, c.active)
}

func TestL4d2_boost_patchServerModule(t *testing.T) {
	mockedWriterFactory := &memory.MockedWriterFactory{}

	buffer := bytes.Buffer{}
	mockedWriterFactory.NewFunc = func(process interface{}, address uintptr) io.Writer {
		return io.Writer(&buffer)
	}

	c := &boostCommand{}
	c.memWriterFactory = mockedWriterFactory

	err := c.patchServerModule(0x99, process.DWORD(123), uintptr(0x123456))
	assert.Nil(t, err)
	assert.Equal(t, []byte{0x99}, buffer.Bytes())
}

func TestL4d2_boost_execServerModuleCode(t *testing.T) {
	runTriggered := false
	mockedThread := &thread.MockedRemote{}
	mockedThread.RunFunc = func() error {
		runTriggered = true
		return nil
	}

	mockedThreadFactory := &thread.MockedRemoteFactory{}
	mockedThreadFactory.NewFunc = func(process interface{}, entryAddress uintptr, argsAddress *uintptr) thread.Remote {
		return mockedThread
	}

	c := &boostCommand{}
	c.remoteThreadFactory = mockedThreadFactory

	err := c.execServerModuleCode(process.DWORD(123), uintptr(0x123456))
	assert.Nil(t, err)
	assert.True(t, runTriggered)
}
