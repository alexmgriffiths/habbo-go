package navigator

import "github.com/alexmgriffiths/habbo-go/communication/outgoing"

func NavigatorMetadataComposer() []byte {
	packet := outgoing.NewOutgoingPacket(3052)
	buffer := packet.GetBuffer()

	buffer.WriteInt(4)
	buffer.WriteString("official_view")
	buffer.WriteInt(0)
	buffer.WriteString("hotel_view")
	buffer.WriteInt(0)
	buffer.WriteString("roomads_view")
	buffer.WriteInt(0)
	buffer.WriteString("myworld_view")
	buffer.WriteInt(0)
	return buffer.Wrap()

}
