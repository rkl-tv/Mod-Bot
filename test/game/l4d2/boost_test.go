package l4d2

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
	"time"
)

func TestL4D2_Boost_success(t *testing.T) {
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

	helper.SendPrivateIrcMessage("666", "rkl85", "RKL85", "$boost")
	time.Sleep(2 * time.Second)
	assert.Contains(t, lBuf.String(), "[R] RKL85 (666): [boost]")
	assert.Contains(t, lBuf.String(), "[I] forward \"[boost]\" request to attached game")
	assert.Contains(t, lBuf.String(), "[I] Initiate team boost for 1 seconds")
	assert.Contains(t, lBuf.String(), "[M] RKL85 (666): @RKL85 your boost has finished")

	messages := helper.GetMessages()
	assert.Contains(t, messages, "@RKL85 your boost has finished")
}

func TestL4D2_Boost_already_active(t *testing.T) {
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

	helper.SendPrivateIrcMessage("666", "rkl85", "RKL85", "$boost")
	helper.SendPrivateIrcMessage("666", "rkl85", "RKL85", "$boost")
	time.Sleep(2 * time.Second)
	assert.Contains(t, lBuf.String(), "[E] boost is already active")
	assert.Contains(t, lBuf.String(), "[M] RKL85 (666): @RKL85 boost is already active")

	messages := helper.GetMessages()
	assert.Contains(t, messages, "@RKL85 boost is already active")
}
