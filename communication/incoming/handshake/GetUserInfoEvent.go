package handshake

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/communication/outgoing/users"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

type GetUserInfoEvent struct{}

func (e *GetUserInfoEvent) Handle(gm *managers.GameManager, packet *incoming.IncomingPacket, client *websocket.Conn) error {

	currentHabbo := gm.GetClient(client)
	responsePacket := users.UserObjectComposer(currentHabbo)

	client.WriteMessage(websocket.BinaryMessage, responsePacket)
	return nil
}
