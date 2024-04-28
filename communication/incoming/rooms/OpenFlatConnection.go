package rooms

import (
	"math/rand"

	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/communication/outgoing/rooms"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

// Initialize the "class"
type OpenFlatConnection struct{}

// Add the expected Handle method to the "class"
func (e *OpenFlatConnection) Handle(managers *managers.Managers, packet *incoming.IncomingPacket, client *websocket.Conn) error {
	buffer := packet.GetBuffer()

	roomId := buffer.ReadInt()

	// TODO: implement
	// password := buffer.ReadString()

	room := managers.GetRoomManager().GetRoom(roomId)
	habbo := managers.GetGameManager().GetClient(client)

	habbo.GetRoomUnit().SetId(rand.Intn(999999) + 100000)
	habbo.GetRoomUnit().SetRoom(room)
	habbo.SetCurrentRoom(room)

	if habbo.GetRoomUnit().GetCurrentLocation() == nil {
		habbo.GetRoomUnit().SetCurrentLocation(room.GetLayout().GetDoorTile())
		if habbo.GetRoomUnit().GetCurrentLocation() != nil {
			habbo.GetRoomUnit().SetZ(habbo.GetRoomUnit().GetCurrentLocation().GetStackHeight())
		}

		habbo.GetRoomUnit().SetBodyRotation(room.GetLayout().GetDoorDir())
		habbo.GetRoomUnit().SetHeadRotation(room.GetLayout().GetDoorDir())
	}

	client.WriteMessage(websocket.BinaryMessage, rooms.OpenConnectionMessageComposer())
	client.WriteMessage(websocket.BinaryMessage, rooms.RoomReadyMessageComposer(room))
	client.WriteMessage(websocket.BinaryMessage, rooms.RoomPropertyMessageComposer("landscape", "0.0"))
	client.WriteMessage(websocket.BinaryMessage, rooms.RoomRatingComposer(room.GetScore(), room.GetOwnerID() != habbo.GetId()))

	return nil
}
