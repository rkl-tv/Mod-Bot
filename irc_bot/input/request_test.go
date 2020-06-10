package input

import (
	"ModBot/twitch/irc"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewRequest(t *testing.T) {
	user := irc.User{ID: "", Name: "", DisplayName: ""}
	raw := "hello world"

	r := NewRequest(user, raw)
	assert.NotNil(t, r)
	assert.True(t, r.createdAt.Before(time.Now().Add(1*time.Second)))
	assert.Equal(t, user, r.ircUser)
	assert.Equal(t, raw, r.raw)
}
