package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDefaultConnector_Attach(t *testing.T) {
	g := &MockedGame{}

	c := NewDefaultConnector().(*DefaultConnector)
	assert.Nil(t, c.attachedGame)

	c.Attach(g)
	assert.Equal(t, g, c.attachedGame)
}

func TestDefaultConnector_GetGame(t *testing.T) {
	g := &MockedGame{}
	c := NewDefaultConnector().(*DefaultConnector)
	c.Attach(g)

	res := c.GetGame()
	assert.Equal(t, g, res)
}
