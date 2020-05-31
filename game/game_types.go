package game

import (
	"net"
)

type player struct {
	con net.Conn
}

// Game : game class
type game struct {
	p1, p2 player
	board  []byte
}
