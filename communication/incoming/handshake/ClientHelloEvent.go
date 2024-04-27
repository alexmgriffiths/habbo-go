package handshake

import (
	"fmt"

	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

type ClientHelloEvent struct{}

func (h *ClientHelloEvent) Handle(gm *managers.GameManager, packet *incoming.IncomingPacket, c *websocket.Conn) error {
	version := packet.GetBuffer().ReadString()
	fmt.Println(version)
	return nil
}
