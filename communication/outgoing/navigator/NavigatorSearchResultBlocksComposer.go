package navigator

import (
	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
)

func NavigatorSearchResultBlocksComposer(view string, query string, rooms map[int32]*rooms.Room) []byte {
	packet := *outgoing.NewOutgoingPacket(2690)
	buffer := packet.GetBuffer()

	buffer.WriteString(view)
	buffer.WriteString(query)

	buffer.WriteInt(1) // ???

	buffer.WriteString(view)
	buffer.WriteString(query)

	buffer.WriteInt(0)
	buffer.WriteBool(false)
	buffer.WriteInt(0) // 0 = List mode

	buffer.WriteInt(int32(len(rooms)))

	for _, room := range rooms {
		buffer.WriteInt(room.GetId())
		buffer.WriteString(room.GetName())
		buffer.WriteInt(room.GetOwnerID())
		buffer.WriteString(room.GetOwnerName())
		buffer.WriteInt(room.GetState())
		buffer.WriteInt(0) // Current user count
		buffer.WriteInt(room.GetUsersMax())
		buffer.WriteString(room.GetDescription())
		buffer.WriteInt(0) // ???
		buffer.WriteInt(room.GetScore())
		buffer.WriteInt(0)
		buffer.WriteInt(room.GetCategory())

		buffer.WriteInt(int32(len(room.GetTags())))
		for _, tag := range room.GetTags() {
			buffer.WriteString(tag)
		}

		buffer.WriteInt(24)
	}

	return buffer.Wrap()
}
