package managers

import (
	"database/sql"
)

type Managers struct {
	database *sql.DB
	rooms    *RoomManager
	game     *GameManager
}

func NewManagers(database *sql.DB, rooms *RoomManager, game *GameManager) *Managers {
	return &Managers{
		database: database,
		rooms:    rooms,
		game:     game,
	}
}

func (m *Managers) GetDatabase() *sql.DB {
	return m.database
}

func (m *Managers) GetRoomManager() *RoomManager {
	return m.rooms
}

func (m *Managers) GetGameManager() *GameManager {
	return m.game
}
