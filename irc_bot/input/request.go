package input

import (
	"ModBot/twitch/irc"
	"time"
)

type Request struct {
	createdAt time.Time
	ircUser   irc.User
	raw       string
}

func NewRequest(ircUser irc.User, raw string) *Request {
	return &Request{
		createdAt: time.Now(),
		ircUser:   ircUser,
		raw:       raw,
	}
}

func (r *Request) GetCreatedAt() time.Time {
	return r.createdAt
}

func (r *Request) GetIrcUser() irc.User {
	return r.ircUser
}

func (r *Request) GetRaw() string {
	return r.raw
}
