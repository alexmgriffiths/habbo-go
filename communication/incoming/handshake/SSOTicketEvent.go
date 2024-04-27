package handshake

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/game"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/alexmgriffiths/habbo-go/utils"
	"github.com/gorilla/websocket"
)

type SSOTicketEvent struct{}

func (h *SSOTicketEvent) Handle(gm *managers.GameManager, packet *incoming.IncomingPacket, conn *websocket.Conn) error {
	ticket := packet.GetBuffer().ReadString()

	outgoing := utils.NewByteBuf([]byte{})
	outgoing.WriteShort(2491)

	habbo, err := game.NewHabbo(gm.GetDatabase(), conn, ticket)

	if err != nil {
		return err
	}

	gm.AddClient(conn, habbo)

	conn.WriteMessage(websocket.BinaryMessage, outgoing.Wrap())

	return nil
}
