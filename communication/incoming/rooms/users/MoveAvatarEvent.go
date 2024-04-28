package users

import (
	"fmt"

	"github.com/alexmgriffiths/habbo-go/communication/incoming"
	"github.com/alexmgriffiths/habbo-go/managers"
	"github.com/gorilla/websocket"
)

type MoveAvatarEvent struct{}

func (e *MoveAvatarEvent) Handle(managers *managers.Managers, packet *incoming.IncomingPacket, client *websocket.Conn) error {

	x := packet.GetBuffer().ReadInt()
	y := packet.GetBuffer().ReadInt()

	habbo := managers.GetGameManager().GetClient(client)
	room := habbo.GetCurrentRoom()
	roomUnit := habbo.GetRoomUnit()

	if roomUnit.GetCurrentLocation().GetX() == int(x) && roomUnit.GetCurrentLocation().GetY() == int(y) {
		fmt.Printf("%d,%d | %d,%d\n", roomUnit.GetCurrentLocation().GetX(), int(x), roomUnit.GetCurrentLocation().GetY(), int(y))
		return nil
	}

	targetTile := room.GetLayout().GetTile(int(x), int(y))
	if targetTile == nil {
		return nil
	}

	roomUnit.SetGoalLocation(targetTile)

	return nil

}
