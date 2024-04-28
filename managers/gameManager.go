package managers

import (
	"database/sql"

	"github.com/alexmgriffiths/habbo-go/game"
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
