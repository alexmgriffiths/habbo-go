package managers

import (
	"github.com/alexmgriffiths/habbo-go/game"
	"github.com/alexmgriffiths/habbo-go/network"
	"github.com/gorilla/websocket"
)

type GameManager struct {
	clients []*network.GameClient
}

func NewGameManager() *GameManager {
	return &GameManager{
		clients: make([]*network.GameClient, 0),
	}
}

func (gm *GameManager) GetClients() []*network.GameClient {
	return gm.clients
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
