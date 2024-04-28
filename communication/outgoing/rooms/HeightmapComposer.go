package rooms

import (
	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
)

func HeightmapComposer(room *rooms.Room) []byte {
	packet := *outgoing.NewOutgoingPacket(2753)
	buffer := packet.GetBuffer()

	buffer.WriteInt(int32(room.GetLayout().GetMapSize() / room.GetLayout().GetMapSizeY()))
	buffer.WriteInt(int32(room.GetLayout().GetMapSize()))

	for y := 0; y < room.GetLayout().GetMapSizeY(); y++ {
		for x := 0; x < room.GetLayout().GetMapSizeX(); x++ {
			t := room.GetLayout().GetTile(x, y)
			if t != nil {
				if t.GetState() == 2 {
					buffer.WriteShort(32767)
				} else {
					buffer.WriteShort(int16(t.GetStackHeight() * 256))
				}
			} else {
				buffer.WriteShort(32767)
			}
		}
	}

	return buffer.Wrap()
}
