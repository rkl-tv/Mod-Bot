package l4d2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestL4D2_Help(t *testing.T) {
	helper := NewBotHelper()

	bot := helper.CreateBot()
	go func() {
		err := bot.Start()
		assert.Nil(t, err)
	}()
	defer func() {
		err := bot.Stop()
		assert.Nil(t, err)
	}()

	var lBuf bytes.Buffer
	log.SetOutput(&lBuf)

	helper.SendPrivateIrcMessage("666", "rkl85", "RKL85", "$help")
	time.Sleep(2 * time.Second)
	assert.Contains(t, lBuf.String(), "[R] RKL85 (666): [help]")
	assert.Contains(t, lBuf.String(), "[I] found native handler for \"[help]\" request")
	assert.Contains(t, lBuf.String(), "[M] RKL85 (666): L4D2 Usage:  \"$help\" -> prints this help,  \"$boost\" -> enables team god mode for 30 seconds")

	messages := helper.GetMessages()
	assert.Contains(t, messages, "L4D2 Usage:  \"$help\" -> prints this help,  \"$boost\" -> enables team god mode for 30 seconds")
}
