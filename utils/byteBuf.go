package utils

import (
	"encoding/binary"
	"errors"
	"io"
)

// ByteBuf is a custom buffer type that implements the io.Reader and io.Writer interfaces
type ByteBuf struct {
	source []byte
	offset int
	length int
}

func (b *ByteBuf) GetSource() []byte {
	return b.source
}

// NewByteBuf creates a new ByteBuf instance
func NewByteBuf(source []byte) *ByteBuf {
	return &ByteBuf{
		source: source,
		offset: 0,
		length: len(source),
	}
}

// ReadByte reads a single byte from the buffer
func (b *ByteBuf) ReadByte() (byte, error) {
	if b.offset >= b.length {
		return 0, io.EOF
	}
	val := b.source[b.offset]
	b.offset++
	return val, nil
}

// ReadBytes reads bytes from the buffer
func (b *ByteBuf) ReadBytes(size int) ([]byte, error) {
	if b.offset+size > b.length {
		return nil, errors.New("not enough data in buffer")
	}
	data := b.source[b.offset : b.offset+size]
	b.offset += size
	return data, nil
}

// ReadInt reads a 32-bit integer from the buffer
func (b *ByteBuf) ReadInt() int32 {
	data, err := b.ReadBytes(4)
	if err != nil {
		return 0
	}
	return int32(binary.BigEndian.Uint32(data))
}

// ReadUInt reads a 32-bit unsigned integer from the buffer
func (b *ByteBuf) ReadUInt() (uint32, error) {
	data, err := b.ReadBytes(4)
	if err != nil {
		return 0, err
	}
	return binary.BigEndian.Uint32(data), nil
}

// ReadShort reads a 16-bit integer from the buffer
func (b *ByteBuf) ReadShort() int16 {
	data, err := b.ReadBytes(2)
	if err != nil {
		return 0
	}
	return int16(binary.BigEndian.Uint16(data))
}

func (b *ByteBuf) ReadString() string {
	strLen := b.ReadShort()
	bytes, err := b.ReadBytes(int(strLen))
	if err != nil {
		return ""
	}
	str := string(bytes)
	return str
}

// WriteByte writes a single byte to the buffer
func (b *ByteBuf) WriteByte(value byte) error {
	b.source = append(b.source, value)
	b.length++
	return nil
}

// WriteBytes writes bytes to the buffer
func (b *ByteBuf) WriteBytes(data []byte) error {
	b.source = append(b.source, data...)
	b.length += len(data)
	return nil
}

// WriteInt writes a 32-bit integer to the buffer
func (b *ByteBuf) WriteInt(value int32) error {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, uint32(value))
	return b.WriteBytes(buf)
}

// WriteUInt writes a 32-bit unsigned integer to the buffer
func (b *ByteBuf) WriteUInt(value uint32) error {
	buf := make([]byte, 4)
	binary.BigEndian.PutUint32(buf, value)
	return b.WriteBytes(buf)
}

// WriteShort writes a 16-bit integer to the buffer
func (b *ByteBuf) WriteShort(value int16) error {
	buf := make([]byte, 2)
	binary.BigEndian.PutUint16(buf, uint16(value))
	return b.WriteBytes(buf)
}

// WriteString writes a string to the buffer
func (b *ByteBuf) WriteString(s string) error {
	err := b.WriteShort(int16(len(s)))
	if err != nil {
		return err
	}
	return b.WriteBytes([]byte(s))
}

// WriteBool writes a boolean to the buffer
func (b *ByteBuf) WriteBool(value bool) error {
	var boolByte byte
	if value {
		boolByte = 1
	} else {
		boolByte = 0
	}
	return b.WriteByte(boolByte)
}

func (b *ByteBuf) Wrap() []byte {
	temp := NewByteBuf([]byte{})
	temp.WriteInt(int32(len(b.GetSource())))
	temp.WriteBytes(b.GetSource())
	return temp.GetSource()
}
