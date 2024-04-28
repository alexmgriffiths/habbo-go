package rooms

import (
	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
)

func RoomReadyMessageComposer(room *rooms.Room) []byte {
	packet := *outgoing.NewOutgoingPacket(2031)
	buffer := packet.GetBuffer()

	buffer.WriteString(room.GetLayout().GetName())
	buffer.WriteInt(room.GetId())

	return buffer.Wrap()
}
