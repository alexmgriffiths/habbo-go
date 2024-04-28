package navigator

import (
	"github.com/alexmgriffiths/habbo-go/communication/outgoing"
	"github.com/alexmgriffiths/habbo-go/game"
	"github.com/alexmgriffiths/habbo-go/managers"
)

func NavigatorEventCategoriesComposer(habbo *game.Habbo, roomManager *managers.RoomManager) []byte {
	packet := *outgoing.NewOutgoingPacket(3244)
	buffer := packet.GetBuffer()

	categories := roomManager.GetCategories()
	buffer.WriteInt(int32(len(categories)))

	for _, category := range categories {
		buffer.WriteInt(int32(category.GetId()))
		buffer.WriteString(category.GetCaption())
		buffer.WriteBool(true)  // Visible
		buffer.WriteBool(false) // ??? true = disconnect
		buffer.WriteString(category.GetCaption())
		buffer.WriteString(category.GetCaption())
		buffer.WriteBool(false)
	}

	return buffer.Wrap()
}
