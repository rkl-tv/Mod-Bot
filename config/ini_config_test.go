package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testIniConfig = &IniConfig{
	getIniPathFunc: func() string {
		return "ini_config.dist.ini"
	},
}

func TestIniConfig_GetTwitchIrcUsername(t *testing.T) {
	res := testIniConfig.GetTwitchIrcUsername()
	assert.Equal(t, "rkl85", res)
}

func TestIniConfig_GetTwitchIrcAuthentication(t *testing.T) {
	res := testIniConfig.GetTwitchIrcAuthentication()
	assert.Equal(t, "oauth:123456", res)
}

func TestIniConfig_GetTwitchIrcChannel(t *testing.T) {
	res := testIniConfig.GetTwitchIrcChannel()
	assert.Equal(t, "rkl85-Channel", res)
}
