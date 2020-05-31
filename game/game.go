package game

import (
	"fmt"
)

func (g game) abort() {
	fmt.Println("game.abort()")
	g.sendAll([]byte("Error"))
}

func (g game) end(winner byte) {
	fmt.Println("game.end() | winner:", winner)
	msg := []byte("GameEnded")
	msg = append(msg, winner)
	g.sendAll(msg)
}

func (g game) start() {
	g.sendAll([]byte("GameStarted"))
	go func() {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("game.recoverErr:", p)
				g.abort()
			}
		}()
		g.board = make([]byte, 9)
		turn, buff, moveCount := byte(1), make([]byte, 32), 0
		activePlayer := g.p1
		for {
			// wait for active player
			_, err := activePlayer.con.Read(buff)
			if err != nil {
				fmt.Println("game.readErr:", err)
				g.abort()
				return
			}
			v := int(buff[0])
			fmt.Println("v:", v)
			if v >= len(g.board) || g.board[v] != 0 {
				for i := 0; i < len(g.board); i++ {
					if g.board[i] == 0 {
						g.board[i] = turn
						break
					}
				}
			} else {
				g.board[v] = 1
			}
			// check for EOG
			// ** change this for dynamic games
			for i := 0; i < 3; i++ {
				r := i * 3
				if g.board[r] != 0 &&
					g.board[r] == g.board[r+1] && g.board[r] == g.board[r+2] {
					g.end(turn)
					return
				}
			}
			for i := 0; i < 3; i++ {
				r := i
				if g.board[r] != 0 &&
					g.board[r] == g.board[r+3] && g.board[r] == g.board[r+6] {
					g.end(turn)
					return
				}
			}
			moveCount++
			if moveCount == len(g.board) {
				g.end(3)
				return
			}
			// change turn
			if turn == 1 {
				activePlayer = g.p2
				turn = 2
			} else {
				activePlayer = g.p1
				turn = 1
			}
		}
	}()
}
