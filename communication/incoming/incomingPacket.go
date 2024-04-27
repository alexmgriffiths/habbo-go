package incoming

import "github.com/alexmgriffiths/habbo-go/utils"

type IncomingPacket struct {
	header uint16
	length int
	buffer utils.ByteBuf
	data   []byte
}

func NewIncomingPacket(header uint16, length int, buffer utils.ByteBuf, data []byte) *IncomingPacket {
	return &IncomingPacket{
		header: header,
		length: length,
		buffer: buffer,
		data:   data,
	}
}

func (ip *IncomingPacket) GetHeader() int {
	return int(ip.header)
}

func (ip *IncomingPacket) GetLength() int {
	return int(ip.length)
}

func (ip *IncomingPacket) GetBuffer() *utils.ByteBuf {
	return &ip.buffer
}

func (ip *IncomingPacket) GetData() []byte {
	return ip.data
}
