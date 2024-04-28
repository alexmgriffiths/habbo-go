package handshake

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/communication/outgoing/handshake"
	"github.com/alexmgriffiths/habbo-go/game"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

type SSOTicketEvent struct{}

func (h *SSOTicketEvent) Handle(managers *managers.Managers, packet *incoming.IncomingPacket, conn *websocket.Conn) error {
	ticket := packet.GetBuffer().ReadString()
	habbo, err := game.NewHabbo(managers.GetDatabase(), conn, ticket)

	if err != nil {
		return err
	}

	managers.GetGameManager().AddClient(conn, habbo)
	conn.WriteMessage(websocket.BinaryMessage, handshake.NewAuthenticationOKMessageComposer())

	return nil
}
