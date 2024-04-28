package managers

import (
	"database/sql"
	"sync"
	"time"

	"github.com/alexmgriffiths/habbo-go/communication/outgoing/users"
	"github.com/alexmgriffiths/habbo-go/game"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
	"github.com/alexmgriffiths/habbo-go/network"
	"github.com/gorilla/websocket"
)

type GameManager struct {
	clients []*network.GameClient
	db      *sql.DB
}

func NewGameManager(db *sql.DB) *GameManager {
	return &GameManager{
		clients: make([]*network.GameClient, 0),
		db:      db,
	}
}

func (gm *GameManager) GetRoomHabbos(roomID int32) []*game.Habbo {
	roomHabbos := []*game.Habbo{}

	habbos := gm.GetClients()
	for _, habbo := range habbos {
		currentRoom := habbo.GetCurrentRoom()
		if currentRoom != nil && currentRoom.GetId() == roomID {
			roomHabbos = append(roomHabbos, habbo)
		}
	}
	return roomHabbos
}

func (gm *GameManager) GetActiveRooms() []*rooms.Room {
	habbos := gm.GetClients()
	rooms := []*rooms.Room{}
	for _, habbo := range habbos {
		if habbo.GetCurrentRoom() != nil {
			found := false
			for _, r := range rooms {
				if r.GetId() == habbo.GetCurrentRoom().GetId() {
					found = true
				}
			}
			if !found {
				rooms = append(rooms, habbo.GetCurrentRoom())
			}
		}
	}
	return rooms
}

func (gm *GameManager) StartCycles() {
	habbos := gm.GetClients()
	if len(habbos) > 0 {
		ticker := time.NewTicker(500 * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			var wg sync.WaitGroup

			for _, room := range gm.GetActiveRooms() {
				wg.Add(1)
				roomHabbos := gm.GetRoomHabbos(room.GetId())
				go func(habbos []*game.Habbo) {
					defer wg.Done()
					for _, h := range habbos {
						// Todo only cycle when needed
						h.GetRoomUnit().Cycle()
						temp := users.UserUpdateComposer(habbos)
						h.GetConnection().WriteMessage(websocket.BinaryMessage, temp)
					}
				}(roomHabbos)
			}
			wg.Wait()
		}
	}
}

func (gm *GameManager) GetDatabase() *sql.DB {
	return gm.db
}

func (gm *GameManager) GetClients() []*game.Habbo {
	var clients []*game.Habbo
	for _, client := range gm.clients {
		clients = append(clients, client.GetHabbo())
	}
	return clients
}

func (gm *GameManager) GetClient(connection *websocket.Conn) *game.Habbo {
	for _, client := range gm.clients {
		if client.GetConnection() == connection {
			return client.GetHabbo()
		}
	}
	return nil
}

func (gm *GameManager) AddClient(client *websocket.Conn, habbo *game.Habbo) {
	gameClient := network.NewGameClient(client, habbo)
	gm.clients = append(gm.clients, gameClient)
}
