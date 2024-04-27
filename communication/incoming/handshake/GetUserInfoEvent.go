package handshake

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/alexmgriffiths/habbo-go/utils"
	"github.com/gorilla/websocket"
)

type GetUserInfoEvent struct{}

func (e *GetUserInfoEvent) Handle(gm *managers.GameManager, packet *incoming.IncomingPacket, client *websocket.Conn) error {

	currentHabbo := gm.GetClient(client)
	response := utils.NewByteBuf([]byte{})
	response.WriteShort(2725)

	response.WriteInt(currentHabbo.GetId())          // User ID
	response.WriteString(currentHabbo.GetUsername()) // Username
	response.WriteString(currentHabbo.GetLook())     // Figure
	response.WriteString(currentHabbo.GetGender())   // Gender
	response.WriteString(currentHabbo.GetMotto())    // Motto
	response.WriteString(currentHabbo.GetUsername()) // Real Name
	response.WriteBool(false)                        // Direct mail
	response.WriteInt(0)                             // respects received
	response.WriteInt(0)                             // respects remaining
	response.WriteInt(0)                             // pet respects remaining
	response.WriteBool(false)                        //StreamPublishedAllowed
	response.WriteString("01-01-1970 00:00:00")      // last access date
	response.WriteBool(false)                        //Can change name
	response.WriteBool(false)                        // Safetly locked

	client.WriteMessage(websocket.BinaryMessage, response.Wrap())
	return nil
}
