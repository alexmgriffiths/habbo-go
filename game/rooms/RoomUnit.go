package rooms

import "fmt"

type RoomUnit struct {
	id              int
	unitType        int
	status          string
	currentLocation *RoomTile
	x               int
	y               int
	z               int
	bodyRotation    int
	headRotation    int

	startLocation    *RoomTile
	previousLocation *RoomTile
	goalLocation     *RoomTile

	room *Room
	path []*RoomTile

	statusUpdate bool
}

func NewRoomUnit() *RoomUnit {

	return &RoomUnit{
		id:       0,
		unitType: 1,
		status:   "/flatctrl/",
	}
}

func (ru *RoomUnit) GetStatus() string {
	return ru.status
}

func (ru *RoomUnit) GetBodyRotation() int32 {
	return int32(ru.bodyRotation)
}

func (ru *RoomUnit) GetHeadRotation() int32 {
	return int32(ru.headRotation)
}

func (ru *RoomUnit) GetCurrentLocation() *RoomTile {
	return ru.currentLocation
}

func (ru *RoomUnit) SetCurrentLocation(tile *RoomTile) {
	ru.currentLocation = tile
}

func (ru *RoomUnit) GetId() int32 {
	return int32(ru.id)
}

func (ru *RoomUnit) SetId(id int) {
	ru.id = id
}

func (ru *RoomUnit) SetRoom(room *Room) {
	ru.room = room
}

func (ru *RoomUnit) SetX(x int) {
	ru.x = x
}

func (ru *RoomUnit) SetY(y int) {
	ru.y = y
}

func (ru *RoomUnit) SetZ(z int) {
	ru.z = z
}

func (ru *RoomUnit) SetBodyRotation(rot int) {
	ru.bodyRotation = rot
}

func (ru *RoomUnit) SetHeadRotation(rot int) {
	ru.headRotation = rot
}

func (ru *RoomUnit) GetStartLocation() *RoomTile {
	return ru.startLocation
}

func (ru *RoomUnit) SetStartLocation(tile *RoomTile) {
	ru.startLocation = tile
}

func (ru *RoomUnit) GetPreviousLocation() *RoomTile {
	return ru.previousLocation
}

func (ru *RoomUnit) SetPreviousLocation(tile *RoomTile) {
	ru.previousLocation = tile
}

func (ru *RoomUnit) GetGoalLocation() *RoomTile {
	return ru.goalLocation
}

func (ru *RoomUnit) SetGoalLocation(tile *RoomTile) {
	if tile != nil {
		ru.startLocation = ru.currentLocation
		ru.goalLocation = tile
		ru.FindPath()
		if len(ru.path) < 1 {
			ru.goalLocation = ru.currentLocation
		}
	} else {
		fmt.Printf("Tile %d,%d is null", tile.GetX(), tile.GetY())
	}
}

func (ru *RoomUnit) FindPath() {
	newPath := ru.GetRoom().GetLayout().FindPath(ru.currentLocation, ru.goalLocation, ru)
	if len(newPath) > 0 {
		ru.path = newPath
	}
}

func (ru *RoomUnit) Cycle() bool {
	if ru.GetId() == 0 {
		return false
	}

	if len(ru.path) > 0 {
		next := ru.path[len(ru.path)-1]
		ru.status = fmt.Sprintf("/mv %d,%d,%d/", next.GetX(), next.GetY(), next.GetZ())
		ru.path = ru.path[:len(ru.path)-1] // array.pop()
		ru.previousLocation = ru.currentLocation
		ru.currentLocation = next
		ru.bodyRotation = ru.CalculateRotation(ru.previousLocation.x, ru.previousLocation.y, next.x, next.y)
		ru.headRotation = ru.CalculateRotation(ru.previousLocation.x, ru.previousLocation.y, next.x, next.y)
		return true
	} else {
		if ru.goalLocation != nil {
			if ru.currentLocation == ru.goalLocation && ru.status == "/flatctrl/" {
				ru.goalLocation = nil
				return false
			}

			ru.currentLocation = ru.goalLocation
			ru.previousLocation = ru.currentLocation
			ru.status = "/flatctrl/"
			return true

		}
	}

	return ru.statusUpdate
}

func (ru *RoomUnit) CalculateRotation(x1 int, y1 int, x2 int, y2 int) int {
	if x1 > x2 && y1 > y2 {
		return 7
	}
	if x1 < x2 && y1 < y2 {
		return 3
	}
	if x1 > x2 && y1 < y2 {
		return 5
	}
	if x1 < x2 && y1 > y2 {
		return 1
	}
	if x1 > x2 {
		return 6
	}
	if x1 < x2 {
		return 2
	}
	if y1 > y2 {
		return 0
	}
	if y1 < y2 {
		return 4
	}
	return 0
}

func (ru *RoomUnit) GetRoom() *Room {
	return ru.room
}
