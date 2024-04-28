package rooms

import (
	"strconv"
	"strings"
)

type RoomLayout struct {
	name      string
	doorX     int
	doorY     int
	doorZ     int
	doorDir   int
	heightmap string

	mapSize  int
	mapSizeX int
	mapSizeY int

	roomTiles     [][]*RoomTile
	highestPoint  int
	squareHeights [][]int
	squareStates  [][]int

	doorTile *RoomTile
}

func NewRoomLayout(name string, doorX int, doorY int, doorDir int, heightmap string) *RoomLayout {
	return &RoomLayout{
		name:      name,
		doorX:     doorX,
		doorY:     doorY,
		doorDir:   doorDir,
		heightmap: heightmap,
	}
}

func (l *RoomLayout) Parse() {
	modelTemp := strings.ReplaceAll(l.heightmap, "\n", "")
	splitModel := strings.Split(modelTemp, "\r")

	l.mapSizeX = len(splitModel[0])
	l.mapSizeY = len(splitModel)

	l.roomTiles = make([][]*RoomTile, l.mapSizeX)
	for i := range l.roomTiles {
		l.roomTiles[i] = make([]*RoomTile, l.mapSizeY)
	}

	l.mapSize = 0

	for y := 0; y < l.mapSizeY; y++ {
		if splitModel[y] == "" || splitModel[y] == "\r" {
			continue
		}

		row := splitModel[y]
		for x := 0; x < l.mapSizeX; x++ {
			if len(row) != l.mapSizeX {
				break
			}
			square := strings.ToLower(strings.TrimSpace(row[x : x+1]))
			state := 0 // 0 = Open
			height := 0

			if square == "x" {
				state = 2
			} else {
				if square == "" {
					height = 0
				} else {
					h, err := strconv.Atoi(square)
					if err == nil {
						height = h
					} else {
						temp := 10 + strings.Index(strings.ToUpper(square), "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
						height = temp
					}
				}
			}

			l.roomTiles[x][y] = NewRoomTile(x, y, height, state, true)
			l.mapSize++

		}

	}

	if len(l.roomTiles) > 0 {
		l.doorTile = l.roomTiles[l.doorX][l.doorY]
		l.doorTile.SetAllowStack(false)

		doorTileFront := l.GetTileInFront(l.doorTile, l.doorDir, 0)
		if doorTileFront != nil && l.TileExists(doorTileFront.GetX(), doorTileFront.GetY()) {
			if l.roomTiles[doorTileFront.GetX()][doorTileFront.GetY()].GetState() != 2 {
				if l.doorZ != l.roomTiles[doorTileFront.GetX()][doorTileFront.GetY()].GetZ() || l.roomTiles[l.doorX][l.doorY].state != l.roomTiles[doorTileFront.GetX()][doorTileFront.GetY()].state {
					l.doorZ = l.roomTiles[doorTileFront.GetX()][doorTileFront.GetY()].GetZ()
					l.roomTiles[l.doorX][l.doorY].state = 1
				}
			}
		}
	}

}

func (l *RoomLayout) GetTileInFront(tile *RoomTile, direction int, offset int) *RoomTile {
	offsetX := 0
	offsetY := 0

	rotation := direction % 8

	switch rotation {
	case 0:
		offsetY--
	case 1:
		offsetX++
		offsetY--
	case 2:
		offsetX++
	case 3:
		offsetX++
		offsetY++
	case 4:
		offsetY++
	case 5:
		offsetX--
		offsetY++
	case 6:
		offsetX--
	case 7:
		offsetX--
		offsetY--
	}

	x := tile.GetX()
	y := tile.GetY()

	for i := 0; i <= offset; i++ {
		x += offsetX
		y += offsetY
	}

	return l.GetTile(x, y)

}

func (l *RoomLayout) GetTile(x int, y int) *RoomTile {
	if l.TileExists(x, y) {
		return l.roomTiles[x][y]
	}
	return nil
}

func (l *RoomLayout) TileExists(x int, y int) bool {
	return !(x < 0 || y < 0 || x >= l.mapSizeX || y >= l.mapSizeY)
}

func (l *RoomLayout) GetName() string {
	return l.name
}

func (l *RoomLayout) GetDoorX() int {
	return l.doorX
}

func (l *RoomLayout) GetDoorY() int {
	return l.doorY
}

func (l *RoomLayout) GetDoorDir() int {
	return l.doorDir
}

func (l *RoomLayout) GetHeightmap() string {
	return l.heightmap
}

func (l *RoomLayout) GetDoorTile() *RoomTile {
	return l.doorTile
}

func (l *RoomLayout) GetMapSize() int {
	return l.mapSize
}

func (l *RoomLayout) GetMapSizeX() int {
	return l.mapSizeX
}

func (l *RoomLayout) GetMapSizeY() int {
	return l.mapSizeY
}

func (l *RoomLayout) GetRelativeHeightmap() string {
	return strings.ReplaceAll(l.heightmap, "\r\n", "\r")
}
