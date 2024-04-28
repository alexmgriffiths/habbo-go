package rooms

import "math"

type RoomTile struct {
	x int
	y int
	z int

	state       int
	stackHeight int
	allowStack  bool

	previous   *RoomTile
	diagonally bool
	gCosts     int
	hCosts     int
}

func NewRoomTile(x int, y int, z int, state int, allowStack bool) *RoomTile {
	return &RoomTile{
		x:          x,
		y:          y,
		z:          z,
		state:      state,
		allowStack: allowStack,
		diagonally: false,
	}
}

func (t *RoomTile) GetPrevious() *RoomTile {
	return t.previous
}

func (t *RoomTile) SetPrevious(tile *RoomTile) {
	t.previous = tile
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

func (t *RoomTile) GetFCosts() int {
	return t.gCosts + t.hCosts
}

func (t *RoomTile) SetGCosts(gCosts int) {
	t.gCosts = gCosts
}

func (t *RoomTile) SetGCostsByBasicCost(previousRoomTile *RoomTile, basicCost int) {
	t.SetGCosts(previousRoomTile.GetGCosts() + basicCost)
}

func (t *RoomTile) SetGCostsByTile(previousRoomTile *RoomTile) {
	if t.diagonally {
		t.SetGCostsByBasicCost(previousRoomTile, 14)
	} else {
		t.SetGCostsByBasicCost(previousRoomTile, 10)
	}
}

func (t *RoomTile) CalculateGCosts(previousRoomTile *RoomTile) int {
	if t.diagonally {
		return previousRoomTile.GetGCosts() + 14
	}
	return previousRoomTile.GetGCosts() + 10
}

func (t *RoomTile) SetHCosts(parent *RoomTile) {

	xDiff := int(math.Abs(float64(t.x) - float64(parent.GetX())))
	yDiff := int(math.Abs(float64(t.y) - float64(parent.GetY())))

	if t.diagonally {
		t.hCosts = int((xDiff + yDiff) * 14)
	} else {
		t.hCosts = int((xDiff + yDiff) * 10)
	}
}

func (t *RoomTile) GetGCosts() int {
	return t.gCosts
}

func (t *RoomTile) SetDiagonally(value bool) {
	t.diagonally = value
}
