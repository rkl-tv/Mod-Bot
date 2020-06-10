package l4d2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewConfig(t *testing.T) {
	c := NewConfig()
	assert.NotNil(t, c)
	assert.Equal(t, uint(30), c.GetBoostSeconds())
}

func TestConfig_SetBoostSeconds(t *testing.T) {
	c := NewConfig()

	c.SetBoostSeconds(uint(1))
	assert.Equal(t, uint(1), c.GetBoostSeconds())
}
