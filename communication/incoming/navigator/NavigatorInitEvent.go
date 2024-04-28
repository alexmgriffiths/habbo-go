package navigator

import (
	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/communication/outgoing/navigator"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

type NavigatorInitEvent struct{}

func (e *NavigatorInitEvent) Handle(managers *managers.Managers, packet *incoming.IncomingPacket, client *websocket.Conn) error {

	habbo := managers.GetGameManager().GetClient(client)
	client.WriteMessage(websocket.BinaryMessage, navigator.NavigatorSettingsComposer(0, 0))
	client.WriteMessage(websocket.BinaryMessage, navigator.NavigatorMetadataComposer())
	client.WriteMessage(websocket.BinaryMessage, navigator.NavigatorLiftedRoomsComposer())
	client.WriteMessage(websocket.BinaryMessage, navigator.NavigatorEventCategoriesComposer(habbo, managers.GetRoomManager()))
	return nil
}
