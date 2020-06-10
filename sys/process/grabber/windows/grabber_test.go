package windows

import (
	error2 "ModBot/sys/error"
	"ModBot/sys/process"
	"github.com/stretchr/testify/assert"
	"testing"
)

const testProcessName = "steam.exe"

func TestNewGrabber(t *testing.T) {
	g := NewGrabber()
	assert.NotNil(t, g)
}

func TestGrabber_Grab(t *testing.T) {
	g := NewGrabber()

	// process not found error
	{
		_, err := g.Grab("deine-mudda.exe")
		assert.IsType(t, &error2.ProcessNotFoundError{}, err)
	}

	// process found
	{
		p, err := g.Grab(testProcessName)
		assert.Nil(t, err)
		assert.Equal(t, testProcessName, p.GetName())
		assert.True(t, p.GetId() > process.DWORD(0))

		modules := p.GetModules()
		assert.Greater(t, len(modules), 100)
	}
}
