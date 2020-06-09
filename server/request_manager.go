package server

import (
	"TicTacToe_Server/game"
	"bufio"
	"fmt"
	"net"
)

func handleConnection(con net.Conn) {
	request, err := bufio.NewReader(con).ReadString('\n')
	if err != nil {
		return
	}
	request = request[:len(request)-1]
	// Log
	fmt.Println("handleConnection/request: ", request, ", len: ", len(request))

	switch request {
	case "Game":
		game.NewGameRequest(con)
	default:
		fmt.Println("request not found")
	}
}
