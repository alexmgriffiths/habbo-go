package handshake

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

type InfoRetrieveEvent struct{}

func (e *InfoRetrieveEvent) Handle(gm *managers.GameManager, packet *incoming.IncomingPacket, client *websocket.Conn) error {

	//habbo := gm.GetClient(client)

	return nil
}
