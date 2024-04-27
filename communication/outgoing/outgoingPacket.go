package outgoing

import "github.com/alexmgriffiths/habbo-go/utils"

type OutgoingPacket struct {
	header int16
	buffer *utils.ByteBuf
}

func NewOutgoingPacket(header int16) *OutgoingPacket {

	buffer := utils.NewByteBuf([]byte{})
	buffer.WriteShort(header)

	return &OutgoingPacket{
		header: header,
		buffer: buffer,
	}
}

func (o *OutgoingPacket) GetBuffer() *utils.ByteBuf {
	return o.buffer
}
