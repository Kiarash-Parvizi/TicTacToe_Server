package game

import (
	"fmt"
	"net"
	"sync"
)

var (
	waitingUser    net.Conn
	hasWaitingUser = false
	mut            sync.Mutex
)

// NewGameRequest : waits for other players for match-making
func NewGameRequest(con net.Conn) {
	fmt.Println("Join Request")
	mut.Lock()
	if waitingUser == nil || !openConnection(waitingUser) {
		hasWaitingUser = false
	}
	if hasWaitingUser {
		// make game
		go tryMakeGame(con, waitingUser)
		hasWaitingUser = false
	} else {
		waitingUser = con
		hasWaitingUser = true
	}
	mut.Unlock()
}

func tryMakeGame(con1, con2 net.Conn) {
	// checks for open connection
	if !openConnection(con1) {
		if openConnection(con2) {
			NewGameRequest(con2)
		}
		return
	}
	if !openConnection(con2) {
		if openConnection(con1) {
			NewGameRequest(con1)
		}
		return
	}
	//
	makeGame(con1, con2)
}

func makeGame(con1, con2 net.Conn) {
	player1 := player{
		con: con1,
	}
	player2 := player{
		con: con2,
	}
	game{
		p1: player1,
		p2: player2,
	}.start()
	fmt.Println("Game was successfully initialized")
}
