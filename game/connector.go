package game

type Connector interface {
	Attach(game Game)
	GetGame() Game
}
