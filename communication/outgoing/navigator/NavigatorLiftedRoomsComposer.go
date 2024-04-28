package navigator

import "github.com/alexmgriffiths/habbo-go/communication/outgoing"

func NavigatorLiftedRoomsComposer() []byte {
	packet := outgoing.NewOutgoingPacket(3104)
	buffer := packet.GetBuffer()

	buffer.WriteInt(0)
	return buffer.Wrap()

}
