package rooms

import (
	"fmt"

	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game"
)

func RoomUserComposer(habbos []*game.Habbo) []byte {
	packet := *outgoing.NewOutgoingPacket(374)
	buffer := packet.GetBuffer()

	buffer.WriteInt(int32(len(habbos)))

	for _, habbo := range habbos {
		buffer.WriteInt(habbo.GetId())
		buffer.WriteString(habbo.GetUsername())
		buffer.WriteString(habbo.GetMotto())
		buffer.WriteString(habbo.GetLook())

		buffer.WriteInt(habbo.GetRoomUnit().GetId())
		buffer.WriteInt(int32(habbo.GetRoomUnit().GetCurrentLocation().GetX()))
		buffer.WriteInt(int32(habbo.GetRoomUnit().GetCurrentLocation().GetY()))

		formattedZ := fmt.Sprintf("%d", habbo.GetRoomUnit().GetCurrentLocation().GetZ())
		buffer.WriteString(formattedZ)
		buffer.WriteInt(habbo.GetRoomUnit().GetBodyRotation())
		buffer.WriteInt(1) // ???
		buffer.WriteString(habbo.GetGender())
		buffer.WriteInt(-1) // Guild
		buffer.WriteInt(-1) // Guild
		buffer.WriteString("")
		buffer.WriteString("")
		buffer.WriteInt(100)
		buffer.WriteBool(true) // ???
	}

	return buffer.Wrap()

}
