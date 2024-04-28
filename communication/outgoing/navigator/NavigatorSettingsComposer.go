package navigator

import "github.com/alexmgriffiths/habbo-go/communication/outgoing"

func NavigatorSettingsComposer(homeroom int32, roomToEnter int32) []byte {
	packet := outgoing.NewOutgoingPacket(518)
	buffer := packet.GetBuffer()
	buffer.WriteInt(homeroom)
	buffer.WriteInt(roomToEnter)
	return buffer.Wrap()
}
