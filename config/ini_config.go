package config

import (
	"fmt"
	"gopkg.in/ini.v1"
	"log"
	"os"
	"runtime"
	"strconv"
)

var GameProcessName string

type IniConfig struct {
	getIniPathFunc func() string
}

func NewIniConfig() Config {
	cfg := &IniConfig{}

	cfg.getIniPathFunc = func() string {
		path := cfg.getUserHomeDir() + string(os.PathSeparator) + "twitch_connector.ini"

		_, err := os.Stat(path)
		if err != nil {
			log.Fatal(fmt.Sprintf("couldn't read config file: %s", path))
		}

		return path
	}

	return cfg
}

func (c *IniConfig) GetGameProcessName() string {
	return GameProcessName
}

func (c *IniConfig) GetTwitchIrcUsername() string {
	return c.readIniKey("twitch", "TwitchIrcUsername")
}

func (c *IniConfig) GetTwitchIrcAuthentication() string {
	return c.readIniKey("twitch", "TwitchIrcAuthentication")
}

func (c *IniConfig) GetTwitchIrcChannel() string {
	return c.readIniKey("twitch", "TwitchIrcChannel")
}

func (c *IniConfig) L4D2GetBoostSeconds() uint {
	strVal := c.readIniKey("l4d2", "BoostSeconds")

	intVal, err := strconv.Atoi(strVal)
	if err != nil {
		panic(err)
	}

	return uint(intVal)
}

func (c *IniConfig) getUserHomeDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return home
	}
	return os.Getenv("HOME")
}

func (c *IniConfig) readIniKey(section, key string) string {
	iniPath := c.getIniPathFunc()
	file, err := ini.Load(iniPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("error loading ini file: %s", iniPath))
	}

	return file.Section(section).Key(key).String()
}
