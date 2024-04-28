package managers

import (
	"database/sql"
	"log"

	outgoingRooms "github.com/alexmgriffiths/habbo-go/communication/outgoing/rooms"
	"github.com/alexmgriffiths/habbo-go/communication/outgoing/users"
	"github.com/alexmgriffiths/habbo-go/game"
	"github.com/alexmgriffiths/habbo-go/game/rooms"
	"github.com/alexmgriffiths/habbo-go/utils"
	"github.com/gorilla/websocket"
)

type RoomManager struct {
	db             *sql.DB
	logger         *utils.Logger
	roomLayouts    map[string]*rooms.RoomLayout
	roomCategories []*rooms.RoomCategory
	rooms          map[int32]*rooms.Room
}

func NewRoomManager(db *sql.DB, logger *utils.Logger) *RoomManager {
	return &RoomManager{
		db:             db,
		roomLayouts:    make(map[string]*rooms.RoomLayout),
		roomCategories: []*rooms.RoomCategory{},     // Empty slice
		rooms:          make(map[int32]*rooms.Room), // Empty map
		logger:         logger,
	}
}

func (rm *RoomManager) LoadLayouts() {
	rows, err := rm.db.Query("SELECT * FROM room_models")
	if err != nil {
		log.Fatal("Failed to load room layouts")
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var (
			name        string
			doorX       int
			doorY       int
			doorDir     int
			heightmap   string
			publicItems sql.NullString
			clubOnly    int = 0
		)

		rows.Scan(&name, &doorX, &doorY, &doorDir, &heightmap, &publicItems, &clubOnly)

		layout := rooms.NewRoomLayout(name, doorX, doorY, doorDir, heightmap)
		layout.Parse()

		count++
		rm.roomLayouts[name] = layout
	}
	rm.logger.Success("Loaded %d room layouts!", count)
}

func (rm *RoomManager) LoadCustomLayouts() {
	rows, err := rm.db.Query("SELECT * FROM room_models_custom")
	if err != nil {
		log.Fatal("Failed to load room layouts")
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		var (
			name        string
			doorX       int
			doorY       int
			doorDir     int
			heightmap   string
			publicItems sql.NullString
			clubOnly    int = 0
		)

		rows.Scan(&name, &doorX, &doorY, &doorDir, &heightmap, &publicItems, &clubOnly)

		layout := rooms.NewRoomLayout(name, doorX, doorY, doorDir, heightmap)
		layout.Parse()

		count++
		rm.roomLayouts[name] = layout
	}
	rm.logger.Success("Loaded %d custom room layouts!", count)
}

func (rm *RoomManager) LoadCategories() {
	rows, err := rm.db.Query("SELECT * FROM navigator_flatcats")
	if err != nil {
		log.Fatal("Failed to get navigator categories")
	}

	defer rows.Close()

	count := 0
	for rows.Next() {
		var (
			id           int
			minRank      int
			captionSave  string
			caption      string
			canTrade     int
			maxUserCount int
			public       int
			listType     int
			orderNum     int
		)

		err := rows.Scan(&id, &minRank, &captionSave, &caption, &canTrade, &maxUserCount, &public, &listType, &orderNum)

		if err != nil {
			return
		}

		roomCategory := rooms.NewRoomCategory(id, minRank, captionSave, caption, canTrade, maxUserCount, public, listType, orderNum)
		rm.roomCategories = append(rm.roomCategories, roomCategory)
		count++

	}
	rm.logger.Success("Loaded %d room categories!", len(rm.roomCategories))
}

func (rm *RoomManager) LoadRooms() {
	rows, err := rm.db.Query("SELECT * FROM rooms")
	if err != nil {
		log.Fatal("Failed to get rooms")
	}
	defer rows.Close()

	count := 0
	roomsHolder := make(map[int32]*rooms.Room)
	for rows.Next() {
		var (
			id                  int
			ownerID             int
			ownerName           string
			name                string
			description         string
			model               string
			password            string
			state               string
			users               int
			usersMax            int
			guildID             int
			category            int
			score               int
			paperFloor          string
			paperWall           string
			paperLandscape      string
			thicknessWall       int
			wallHeight          int
			thicknessFloor      int
			moodlightData       string
			tags                string
			isPublic            string
			isStaffPicked       string
			allowOtherPets      string
			allowOtherPetsEat   string
			allowWalkthrough    string
			allowHidewall       string
			chatMode            int
			chatWeight          int
			chatSpeed           int
			chatHearingDistance int
			chatProtection      int
			overrideModel       string
			whoCanMute          int
			whoCanKick          int
			whoCanBan           int
			pollID              int
			rollerSpeed         int
			promoted            string
			tradeMode           int
			moveDiagonally      string
			jukeboxActive       string
			hidewired           string
			isForsale           string
		)

		err := rows.Scan(&id, &ownerID, &ownerName, &name, &description, &model, &password, &state, &users, &usersMax, &guildID, &category, &score, &paperFloor, &paperWall, &paperLandscape, &thicknessWall, &wallHeight, &thicknessFloor, &moodlightData, &tags, &isPublic, &isStaffPicked, &allowOtherPets, &allowOtherPetsEat, &allowWalkthrough, &allowHidewall, &chatMode, &chatWeight, &chatSpeed, &chatHearingDistance, &chatProtection, &overrideModel, &whoCanMute, &whoCanKick, &whoCanBan, &pollID, &rollerSpeed, &promoted, &tradeMode, &moveDiagonally, &jukeboxActive, &hidewired, &isForsale)

		if err != nil {
			return
		}
		room := rooms.NewRoom(id, ownerID, ownerName, name, description, model, password, state, users, usersMax, guildID, category, score, paperFloor, paperWall, paperLandscape, thicknessWall, wallHeight, thicknessFloor, moodlightData, tags, isPublic, isStaffPicked, allowOtherPets, allowOtherPetsEat, allowWalkthrough, allowHidewall, chatMode, chatWeight, chatSpeed, chatHearingDistance, chatProtection, overrideModel, whoCanMute, whoCanKick, whoCanBan, pollID, rollerSpeed, promoted, tradeMode, moveDiagonally, jukeboxActive, hidewired, isForsale, rm.roomLayouts[model])

		roomsHolder[int32(id)] = room
		count++
	}

	rm.rooms = roomsHolder
	rm.logger.Success("Loaded %d Rooms!", count)
}

func (rm *RoomManager) GetRooms() map[int32]*rooms.Room {
	return rm.rooms
}

func (rm *RoomManager) AddRoom(room *rooms.Room) {
	rm.rooms[int32(room.GetId())] = room
}

func (rm *RoomManager) GetRoom(roomId int32) *rooms.Room {
	room := rm.rooms[roomId]
	if room == nil {
		row := rm.db.QueryRow("SELECT * FROM rooms WHERE id = ?", roomId)

		var (
			id                  int
			ownerID             int
			ownerName           string
			name                string
			description         string
			model               string
			password            string
			state               string
			users               int
			usersMax            int
			guildID             int
			category            int
			score               int
			paperFloor          string
			paperWall           string
			paperLandscape      string
			thicknessWall       int
			wallHeight          int
			thicknessFloor      int
			moodlightData       string
			tags                string
			isPublic            string
			isStaffPicked       string
			allowOtherPets      string
			allowOtherPetsEat   string
			allowWalkthrough    string
			allowHidewall       string
			chatMode            int
			chatWeight          int
			chatSpeed           int
			chatHearingDistance int
			chatProtection      int
			overrideModel       string
			whoCanMute          int
			whoCanKick          int
			whoCanBan           int
			pollID              int
			rollerSpeed         int
			promoted            string
			tradeMode           int
			moveDiagonally      string
			jukeboxActive       string
			hidewired           string
			isForsale           string
		)

		err := row.Scan(&id, &ownerID, &ownerName, &name, &description, &model, &password, &state, &users, &usersMax, &guildID, &category, &score, &paperFloor, &paperWall, &paperLandscape, &thicknessWall, &wallHeight, &thicknessFloor, &moodlightData, &tags, &isPublic, &isStaffPicked, &allowOtherPets, &allowOtherPetsEat, &allowWalkthrough, &allowHidewall, &chatMode, &chatWeight, &chatSpeed, &chatHearingDistance, &chatProtection, &overrideModel, &whoCanMute, &whoCanKick, &whoCanBan, &pollID, &rollerSpeed, &promoted, &tradeMode, &moveDiagonally, &jukeboxActive, &hidewired, &isForsale)

		if err != nil {
			return nil
		}

		room := rooms.NewRoom(id, ownerID, ownerName, name, description, model, password, state, users, usersMax, guildID, category, score, paperFloor, paperWall, paperLandscape, thicknessWall, wallHeight, thicknessFloor, moodlightData, tags, isPublic, isStaffPicked, allowOtherPets, allowOtherPetsEat, allowWalkthrough, allowHidewall, chatMode, chatWeight, chatSpeed, chatHearingDistance, chatProtection, overrideModel, whoCanMute, whoCanKick, whoCanBan, pollID, rollerSpeed, promoted, tradeMode, moveDiagonally, jukeboxActive, hidewired, isForsale, rm.roomLayouts[model])

		rm.AddRoom(room)
		return room
	}
	return room
}

func (rm *RoomManager) GetCategories() []*rooms.RoomCategory {
	return rm.roomCategories
}

func (rm *RoomManager) GetLayouts() map[string]*rooms.RoomLayout {
	return rm.roomLayouts
}

func (rm *RoomManager) EnterRoom(gm *GameManager, habbo *game.Habbo, room *rooms.Room) {
	if room == nil {
		return
	}

	// if habbo.GetRoomUnit() != nil {
	// 	rm.logger.Error("Users room unit is not null! Already in a room!!!!")
	//	Send userRemoveMessageComposer to room
	// 	return
	// }

	// This might be redunant since we check for this in other places
	// Look into why this has to run every time but doesn't in the node version
	//if habbo.GetRoomUnit().GetCurrentLocation() == nil {
	doorTile := room.GetLayout().GetTile(room.GetLayout().GetDoorX(), room.GetLayout().GetDoorY())
	habbo.GetRoomUnit().SetCurrentLocation(doorTile)
	habbo.GetRoomUnit().GetCurrentLocation().SetZ(doorTile.GetStackHeight())
	habbo.GetRoomUnit().SetBodyRotation(room.GetLayout().GetDoorDir())
	habbo.GetRoomUnit().SetHeadRotation(room.GetLayout().GetDoorDir())
	//}

	// push to array containing room habbos
	roomHabbos := gm.GetRoomHabbos(room.GetId())
	roomUserCount := len(roomHabbos)
	if roomUserCount > 0 {
		for _, gameHabbo := range roomHabbos {
			gameHabbo.GetConnection().WriteMessage(websocket.BinaryMessage, outgoingRooms.RoomUserComposer(roomHabbos))
			gameHabbo.GetConnection().WriteMessage(websocket.BinaryMessage, users.UserUpdateComposer(roomHabbos))
		}
	}

	habbo.GetConnection().WriteMessage(websocket.BinaryMessage, outgoingRooms.RoomEntryInfoComposer(room))
	habbo.GetConnection().WriteMessage(websocket.BinaryMessage, outgoingRooms.RoomVisualizationSettingsComposer(room))
	habbo.GetConnection().WriteMessage(websocket.BinaryMessage, outgoingRooms.GetGuestRoomResultComposer(room, roomUserCount, false, true))

}
