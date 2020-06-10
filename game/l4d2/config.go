package l4d2

type Config struct {
	boostSeconds uint
}

func NewConfig() *Config {
	return &Config{
		boostSeconds: uint(30),
	}
}

func (c *Config) GetBoostSeconds() uint {
	return c.boostSeconds
}

func (c *Config) SetBoostSeconds(s uint) {
	c.boostSeconds = s
}
