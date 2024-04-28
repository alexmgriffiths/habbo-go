package rooms

import (
	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
)

func GetGuestRoomResultComposer(room *rooms.Room, userCount int, roomForward bool, enterRoom bool) []byte {

	packet := *outgoing.NewOutgoingPacket(687)
	buffer := packet.GetBuffer()

	buffer.WriteBool(enterRoom)
	buffer.WriteInt(room.GetId())
	buffer.WriteString(room.GetName())
	buffer.WriteInt(room.GetOwnerID())
	buffer.WriteString(room.GetOwnerName())
	buffer.WriteInt(room.GetState())
	buffer.WriteInt(int32(userCount))
	buffer.WriteInt(room.GetUsersMax())
	buffer.WriteString(room.GetDescription())
	buffer.WriteInt(0)
	buffer.WriteInt(room.GetScore())
	buffer.WriteInt(2)
	buffer.WriteInt(room.GetCategory())

	buffer.WriteInt(0)
	for _, t := range room.GetTags() {
		buffer.WriteString(t)
	}

	buffer.WriteInt(24)

	buffer.WriteBool(roomForward)
	buffer.WriteBool(room.GetStaffPicked())
	buffer.WriteBool(false)
	buffer.WriteBool(false)

	buffer.WriteInt(room.GetMuteOption())
	buffer.WriteInt(room.GetKickOption())
	buffer.WriteInt(room.GetBanOption())

	buffer.WriteBool(true)

	buffer.WriteInt(room.GetChatMode())
	buffer.WriteInt(room.GetChatWeight())
	buffer.WriteInt(room.GetChatSpeed())
	buffer.WriteInt(room.GetChatDistance())
	buffer.WriteInt(room.GetChatProtection())

	return buffer.Wrap()

}
