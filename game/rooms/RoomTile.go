package rooms

type RoomTile struct {
	x int
	y int
	z int

	state       int
	stackHeight int
	allowStack  bool
}

func NewRoomTile(x int, y int, z int, state int, allowStack bool) *RoomTile {
	return &RoomTile{
		x:          x,
		y:          y,
		z:          z,
		state:      state,
		allowStack: allowStack,
	}
}

func (t *RoomTile) GetX() int {
	return t.x
}

func (t *RoomTile) GetY() int {
	return t.y
}

func (t *RoomTile) GetZ() int {
	return t.z
}

func (t *RoomTile) SetX(x int) {
	t.x = x
}

func (t *RoomTile) SetY(y int) {
	t.y = y
}

func (t *RoomTile) SetZ(z int) {
	t.z = z
}

func (t *RoomTile) GetState() int {
	return t.state
}

func (t *RoomTile) GetStackHeight() int {
	return t.stackHeight
}

func (t *RoomTile) SetAllowStack(value bool) {
	t.allowStack = value
}
