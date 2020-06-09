package game

import (
	"net"
	"time"
)

// OpenConnection : checks whether the connection is open or not
func openConnection(c net.Conn) bool {
	oneBuf := make([]byte, 1)
	c.SetReadDeadline(time.Now().Add(time.Second / 10))
	_, err := c.Read(oneBuf)
	// if it's not a timeout error
	if err, ok := err.(net.Error); ok && !err.Timeout() {
		c.Close()
		return false
	}
	c.SetReadDeadline(time.Time{})
	return true
}

func (g *game) sendAll(msgType string, buf []byte) {
	bits := []byte(msgType + "|")
	bits = append(bits, buf...)
	bits = append(bits, '\n')
	g.p1.con.Write(bits)
	g.p2.con.Write(bits)
}

func (g *game) send1(msgType string, buf []byte) {
	bits := []byte(msgType + "|")
	bits = append(bits, buf...)
	bits = append(bits, '\n')
	g.p1.con.Write(bits)
}
func (g *game) send2(msgType string, buf []byte) {
	bits := []byte(msgType + "|")
	bits = append(bits, buf...)
	bits = append(bits, '\n')
	g.p2.con.Write(bits)
}
