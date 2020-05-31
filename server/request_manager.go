package server

import (
	"TicTacToe_Server/game"
	"fmt"
	"net"
)

func handleConnection(con net.Conn) {
	buffer := make([]byte, 64)
	n, err := con.Read(buffer)
	if err != nil {
		return
	}
	request := string(buffer[:n])
	// Log
	fmt.Println("handleConnection/request: ", request, "len: ", len(request))

	switch request {
	case "Game":
		game.NewGameRequest(con)
	default:
		fmt.Println("request not found")
	}
}
