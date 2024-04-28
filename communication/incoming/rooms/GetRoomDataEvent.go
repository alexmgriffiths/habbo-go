package rooms

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/communication/outgoing/rooms"
	"github.com/alexmgriffiths/habbo-go/managers"

	"github.com/gorilla/websocket"
)

type GetRoomDataEvent struct{}

func (e *GetRoomDataEvent) Handle(managers *managers.Managers, packet *incoming.IncomingPacket, client *websocket.Conn) error {
	habbo := managers.GetGameManager().GetClient(client)
	room := habbo.GetCurrentRoom()

	client.WriteMessage(websocket.BinaryMessage, rooms.HeightmapComposer(room))
	client.WriteMessage(websocket.BinaryMessage, rooms.FloorHeightmapComposer(room))

	managers.GetRoomManager().EnterRoom(managers.GetGameManager(), habbo, room)

	return nil
}
