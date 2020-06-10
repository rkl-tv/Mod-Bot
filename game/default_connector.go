package game

type DefaultConnector struct {
	attachedGame Game
}

func NewDefaultConnector() Connector {
	return &DefaultConnector{}
}

func (c *DefaultConnector) Attach(game Game) {
	c.attachedGame = game
}

func (c *DefaultConnector) GetGame() Game {
	return c.attachedGame
}
