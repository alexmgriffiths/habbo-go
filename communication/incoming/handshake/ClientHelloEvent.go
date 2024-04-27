package handshake

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

type ClientHelloEvent struct{}

// This has to be a pointer with context so that we change the struct to have this function, making ClientHelloEvent.Handle accessible
func (h *ClientHelloEvent) Handle(gm *managers.GameManager, packet *incoming.IncomingPacket, c *websocket.Conn) error {
	packet.GetBuffer().ReadString() // Unused, just shows version (NITRO-1-6-6)
	return nil
}
