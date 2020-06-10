package config

type Config interface {
	GetGameProcessName() string
	GetTwitchIrcUsername() string
	GetTwitchIrcAuthentication() string
	GetTwitchIrcChannel() string
	L4D2GetBoostSeconds() uint
}
