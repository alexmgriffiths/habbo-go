package rooms

import (
	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
)

func RoomEntryInfoComposer(room *rooms.Room) []byte {
	packet := *outgoing.NewOutgoingPacket(-1)
	buffer := packet.GetBuffer()

	buffer.WriteInt(room.GetOwnerID())
	buffer.WriteString(room.GetOwnerName())
	return buffer.Wrap()
}
