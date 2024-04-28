package rooms

import "github.com/alexmgriffiths/habbo-go/communication/outgoing"

func RoomPropertyMessageComposer(key string, value string) []byte {
	packet := *outgoing.NewOutgoingPacket(2454)
	buffer := packet.GetBuffer()

	buffer.WriteString(key)
	buffer.WriteString(value)

	return buffer.Wrap()
}
