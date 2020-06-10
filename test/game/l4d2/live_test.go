package l4d2

import (
	"ModBot/game/l4d2"
	windows2 "ModBot/sys/memory/windows"
	"ModBot/sys/process/grabber/windows"
	windows3 "ModBot/sys/process/thread/windows"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBoost(t *testing.T) {
	t.Skip("To use this test, you need a running L4D2 process.")

	cfg := l4d2.NewConfig()
	cfg.SetBoostSeconds(1)

	grabber := windows.NewGrabber()
	game := l4d2.NewL4D2(grabber, windows2.NewMemWriterFactory(), windows3.NewRemoteThreadFactory(), cfg)

	res, err := game.ProcessRequest([]string{"boost"})
	assert.Nil(t, err)
	assert.Contains(t, res.GetMessage(), "your boost has finished")
}
