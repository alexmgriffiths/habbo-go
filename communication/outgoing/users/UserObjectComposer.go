package users

import (
	"fmt"

	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game"
)

func UserObjectComposer(habbo *game.Habbo) []byte {
	packet := outgoing.NewOutgoingPacket(2725)
	buffer := packet.GetBuffer()

	buffer.WriteInt(habbo.GetId())
	buffer.WriteString(habbo.GetUsername())
	buffer.WriteString(habbo.GetLook())
	buffer.WriteString(habbo.GetGender())
	buffer.WriteString(habbo.GetMotto())
	buffer.WriteString(habbo.GetUsername())   // real name
	buffer.WriteBool(false)                   // direct mail
	buffer.WriteInt(0)                        // respects received
	buffer.WriteInt(0)                        // respects to send
	buffer.WriteInt(0)                        // daily pet respects
	buffer.WriteBool(false)                   // stream publish allowed
	buffer.WriteString("01-01-1970 00:00:00") // Last active
	buffer.WriteBool(false)                   // can change name
	buffer.WriteBool(false)                   // Safety locked

	fmt.Println(habbo.GetLook())

	return buffer.Wrap()

}
