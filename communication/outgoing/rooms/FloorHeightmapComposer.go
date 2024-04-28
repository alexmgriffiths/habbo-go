package rooms

import (
	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
)

func FloorHeightmapComposer(room *rooms.Room) []byte {
	packet := *outgoing.NewOutgoingPacket(1301)
	buffer := packet.GetBuffer()

	buffer.WriteBool(true)
	buffer.WriteInt(room.GetWallHeight())
	buffer.WriteString(room.GetLayout().GetRelativeHeightmap())

	return buffer.Wrap()
}
