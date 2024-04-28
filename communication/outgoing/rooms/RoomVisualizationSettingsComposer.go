package rooms

import (
	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
)

func RoomVisualizationSettingsComposer(room *rooms.Room) []byte {
	packet := *outgoing.NewOutgoingPacket(3547)
	buffer := packet.GetBuffer()

	buffer.WriteBool(room.GetHideWall())
	buffer.WriteInt(room.GetWallSize())
	buffer.WriteInt(room.GetFloorSize())

	return buffer.Wrap()

}
