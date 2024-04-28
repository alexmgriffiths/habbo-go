package rooms

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
