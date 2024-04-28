package rooms

import "github.com/alexmgriffiths/habbo-go/communication/outgoing"

func RoomRatingComposer(score int32, canVote bool) []byte {
	packet := *outgoing.NewOutgoingPacket(482)
	buffer := packet.GetBuffer()

	buffer.WriteInt(score)
	buffer.WriteBool(canVote)

	return buffer.Wrap()
}
