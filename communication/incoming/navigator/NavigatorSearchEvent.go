package navigator

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/communication/outgoing/navigator"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

type NavigatorSearchEvent struct{}

func (e *NavigatorSearchEvent) Handle(managers *managers.Managers, packet *incoming.IncomingPacket, client *websocket.Conn) error {
	buffer := packet.GetBuffer()
	view := buffer.ReadString()
	query := buffer.ReadString()

	if view == "query" {
		view = "hotel_view"
	}

	if query == "groups" {
		view = "hotel_view"
	}

	rooms := managers.GetRoomManager().GetRooms()
	client.WriteMessage(websocket.BinaryMessage, navigator.NavigatorSearchResultBlocksComposer(view, query, rooms))
	return nil
}
