package network

import (
	"github.com/alexmgriffiths/habbo-go/game"
	"github.com/gorilla/websocket"
)

type GameClient struct {
	connection *websocket.Conn
	habbo      *game.Habbo
}

func NewGameClient(conn *websocket.Conn, hab *game.Habbo) *GameClient {
	return &GameClient{
		connection: conn,
		habbo:      hab,
	}
}

func (gm *GameClient) GetConnection() *websocket.Conn {
	return gm.connection
}

func (gm *GameClient) GetHabbo() *game.Habbo {
	return gm.habbo
}
