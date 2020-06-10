package l4d2

import (
	error2 "ModBot/game/l4d2/error"
	"ModBot/sys/memory/windows"
	grabber2 "ModBot/sys/process/grabber"
	windows2 "ModBot/sys/process/thread/windows"
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestNewL4D2(t *testing.T) {
	grabber := &grabber2.MockedGrabber{}
	mwFac := windows.NewMemWriterFactory()
	rtFac := windows2.NewRemoteThreadFactory()

	g := NewL4D2(grabber, mwFac, rtFac, NewConfig()).(*l4d2)
	assert.Equal(t, processName, g.processName)
	assert.Equal(t, grabber, g.processGrabber)
	assert.Equal(t, mwFac, g.memWriterFactory)
	assert.Equal(t, rtFac, g.remoteThreadFactory)
}

func TestL4d2_GetUsage(t *testing.T) {
	g := NewL4D2(&grabber2.MockedGrabber{}, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory(), NewConfig())

	usage := g.GetUsage()
	assert.Equal(t, l4d2Usage, usage)
}

func TestL4d2_ProcessRequest(t *testing.T) {
	// "boost"
	{
		// success
		{
			g := NewL4D2(&grabber2.MockedGrabber{}, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory(), NewConfig())

			var lBuf bytes.Buffer
			log.SetOutput(&lBuf)

			_, _ = g.ProcessRequest([]string{"boost"})
			assert.Contains(t, lBuf.String(), "[I] Initiate team boost for 30 seconds\n")
		}

		// already running
		{
			g := NewL4D2(&grabber2.MockedGrabber{}, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory(), NewConfig()).(*l4d2)
			g.boostCommand.active = true

			_, err := g.ProcessRequest([]string{"boost"})
			assert.IsType(t, &error2.BoostIsActiveError{}, err)
		}
	}

	// show usage if command is unknown
	{
		g := NewL4D2(&grabber2.MockedGrabber{}, windows.NewMemWriterFactory(), windows2.NewRemoteThreadFactory(), NewConfig())

		res, err := g.ProcessRequest([]string{"unknown"})
		assert.Nil(t, err)
		assert.Contains(t, res.GetMessage(), "L4D2 Usage:")
	}
}
