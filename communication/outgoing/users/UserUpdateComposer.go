package users

import (
	"fmt"

	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game"
)

func UserUpdateComposer(habbos []*game.Habbo) []byte {
	packet := outgoing.NewOutgoingPacket(1640)
	buffer := packet.GetBuffer()

	buffer.WriteInt(int32(len(habbos)))
	for _, habbo := range habbos {
		if habbo.GetRoomUnit().GetPreviousLocation() != nil {
			buffer.WriteInt(habbo.GetRoomUnit().GetId())
			buffer.WriteInt(int32(habbo.GetRoomUnit().GetPreviousLocation().GetX()))
			buffer.WriteInt(int32(habbo.GetRoomUnit().GetPreviousLocation().GetY()))
			buffer.WriteString(fmt.Sprintf("%d", habbo.GetRoomUnit().GetPreviousLocation().GetZ()))
			buffer.WriteInt(int32(habbo.GetRoomUnit().GetHeadRotation()))
			buffer.WriteInt(int32(habbo.GetRoomUnit().GetBodyRotation()))
			buffer.WriteString(habbo.GetRoomUnit().GetStatus())
		} else {
			buffer.WriteInt(habbo.GetRoomUnit().GetId())
			buffer.WriteInt(int32(habbo.GetRoomUnit().GetCurrentLocation().GetX()))
			buffer.WriteInt(int32(habbo.GetRoomUnit().GetCurrentLocation().GetY()))
			formattedZ := fmt.Sprintf("%d", habbo.GetRoomUnit().GetCurrentLocation().GetZ())
			buffer.WriteString(formattedZ)
			buffer.WriteInt(int32(habbo.GetRoomUnit().GetHeadRotation()))
			buffer.WriteInt(int32(habbo.GetRoomUnit().GetBodyRotation()))
			buffer.WriteString(habbo.GetRoomUnit().GetStatus())
		}
	}

	return buffer.Wrap()
}
