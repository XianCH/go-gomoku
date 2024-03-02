package game

import "net"

type Player struct {
	name string
	ID   int
	conn net.Conn
}

func NewPlayer(name string, ID int) *Player {
	return &Player{
		name: name,
		ID:   ID,
	}
}
