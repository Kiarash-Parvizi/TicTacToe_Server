package game

import (
	"bufio"
	"fmt"
)

func (g game) abort() {
	fmt.Println("game.abort()")
	g.sendAll("abort", []byte("Error"))
}

func (g game) end(winner byte) {
	fmt.Println("game.end() | winner:", winner)
	msg := make([]byte, 1)
	msg[0] = winner
	g.sendAll("end", msg)
}

func (g game) start() {
	g.send1("turn", []byte("1"))
	g.send2("turn", []byte("0"))
	go func() {
		defer func() {
			if p := recover(); p != nil {
				fmt.Println("game.recoverErr:", p)
				g.abort()
			}
		}()
		g.board = make([]byte, 9)
		turn, moveCount := byte(1), 0
		ioChan1, ioChan2 := bufio.NewReader(g.p1.con),
			bufio.NewReader(g.p2.con)
		activePlayer := ioChan1
		for {
			// wait for active player
			req, err := activePlayer.ReadBytes('\n')
			if err != nil {
				fmt.Println("game.readErr:", err)
				g.abort()
				return
			}
			v := int(req[0])
			fmt.Println("v:", v)
			if v >= len(g.board) || g.board[v] != 0 {
				for i := 0; i < len(g.board); i++ {
					if g.board[i] == 0 {
						g.board[i] = turn
						break
					}
				}
			} else {
				g.board[v] = turn
			}
			// check for EOG
			// ** change this for dynamic games
			for i := 0; i < 3; i++ {
				r := i * 3
				if g.board[r] != 0 &&
					g.board[r] == g.board[r+1] && g.board[r] == g.board[r+2] {
					fmt.Println("case 1")
					g.end(turn)
					return
				}
			}
			for i := 0; i < 3; i++ {
				r := i
				if g.board[r] != 0 &&
					g.board[r] == g.board[r+3] && g.board[r] == g.board[r+6] {
					fmt.Println("case 2")
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
				activePlayer = ioChan2
				turn = 2
			} else {
				activePlayer = ioChan1
				turn = 1
			}
			// send state
			g.sendAll("state", g.board)
		}
	}()
}
