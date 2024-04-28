package rooms

import "github.com/alexmgriffiths/habbo-go/communication/outgoing"

func OpenConnectionMessageComposer() []byte {
	packet := *outgoing.NewOutgoingPacket(758)
	return packet.GetBuffer().Wrap()
}
