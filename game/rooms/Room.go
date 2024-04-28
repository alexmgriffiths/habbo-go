package rooms

import (
	"strings"
)

type Room struct {
	ID                  int
	OwnerID             int
	OwnerName           string
	Name                string
	Description         string
	Model               string
	Password            string
	State               string
	Users               int
	UsersMax            int
	GuildID             int
	Category            int
	Score               int
	PaperFloor          string
	PaperWall           string
	PaperLandscape      string
	ThicknessWall       int
	WallHeight          int
	ThicknessFloor      int
	MoodlightData       string
	Tags                string
	IsPublic            string
	IsStaffPicked       string
	AllowOtherPets      string
	AllowOtherPetsEat   string
	AllowWalkthrough    string
	AllowHidewall       string
	ChatMode            int
	ChatWeight          int
	ChatSpeed           int
	ChatHearingDistance int
	ChatProtection      int
	OverrideModel       string
	WhoCanMute          int
	WhoCanKick          int
	WhoCanBan           int
	PollID              int
	RollerSpeed         int
	Promoted            string
	TradeMode           int
	MoveDiagonally      string
	JukeboxActive       string
	Hidewired           string
	IsForsale           string
	layout              *RoomLayout
}

func NewRoom(id int, ownerID int, ownerName string, name string, description string, model string, password string, state string, users int, usersMax int, guildID int, category int, score int, paperFloor string, paperWall string, paperLandscape string, thicknessWall int, wallHeight int, thicknessFloor int, moodlightData string, tags string, isPublic string, isStaffPicked string, allowOtherPets string, allowOtherPetsEat string, allowWalkthrough string, allowHidewall string, chatMode int, chatWeight int, chatSpeed int, chatHearingDistance int, chatProtection int, overrideModel string, whoCanMute int, whoCanKick int, whoCanBan int, pollID int, rollerSpeed int, promoted string, tradeMode int, moveDiagonally string, jukeboxActive string, hidewired string, isForsale string, layout *RoomLayout) *Room {

	return &Room{
		ID:                  id,
		OwnerID:             ownerID,
		OwnerName:           ownerName,
		Name:                name,
		Description:         description,
		Model:               model,
		Password:            password,
		State:               state,
		Users:               users,
		UsersMax:            usersMax,
		GuildID:             guildID,
		Category:            category,
		Score:               score,
		PaperFloor:          paperFloor,
		PaperWall:           paperWall,
		PaperLandscape:      paperLandscape,
		ThicknessWall:       thicknessWall,
		WallHeight:          wallHeight,
		ThicknessFloor:      thicknessFloor,
		MoodlightData:       moodlightData,
		Tags:                tags,
		IsPublic:            isPublic,
		IsStaffPicked:       isStaffPicked,
		AllowOtherPets:      allowOtherPets,
		AllowOtherPetsEat:   allowOtherPetsEat,
		AllowWalkthrough:    allowWalkthrough,
		AllowHidewall:       allowHidewall,
		ChatMode:            chatMode,
		ChatWeight:          chatWeight,
		ChatSpeed:           chatSpeed,
		ChatHearingDistance: chatHearingDistance,
		ChatProtection:      chatProtection,
		OverrideModel:       overrideModel,
		WhoCanMute:          whoCanMute,
		WhoCanKick:          whoCanKick,
		WhoCanBan:           whoCanBan,
		PollID:              pollID,
		RollerSpeed:         rollerSpeed,
		Promoted:            promoted,
		TradeMode:           tradeMode,
		MoveDiagonally:      moveDiagonally,
		JukeboxActive:       jukeboxActive,
		Hidewired:           hidewired,
		IsForsale:           isForsale,
		layout:              layout,
	}
}

func (r *Room) GetWallHeight() int32 {
	return int32(r.WallHeight)
}

// func (r *Room) LoadLayout(db *sql.DB) {
// 	var (
// 		id          int
// 		name        string
// 		doorX       int
// 		doorY       int
// 		doorDir     int
// 		heightmap   string
// 		publicItems sql.NullString
// 		clubOnly    int = 0
// 	)

// 	if r.OverrideModel != "" {
// 		queryRow := db.QueryRow("SELECT * ROM room_models_custom WHERE id = ? LIMIT 1", r.ID)
// 		queryRow.Scan(&id, &name, &doorX, &doorY, &doorDir, &heightmap)

// 	} else {
// 		queryRow := db.QueryRow("SELECT * FROM room_models WHERE name = ? LIMIT 1", r.Model)
// 		queryRow.Scan(&name, &doorX, &doorY, &doorDir, &heightmap, &publicItems, &clubOnly)
// 	}

// 	layout := NewRoomLayout(name, doorX, doorY, doorDir, heightmap)
// 	layout.Parse()

// 	r.layout = layout

// }

func (r *Room) GetId() int32 {
	return int32(r.ID)
}

func (r *Room) GetName() string {
	return r.Name
}

func (r *Room) GetOwnerID() int32 {
	return int32(r.OwnerID)
}

func (r *Room) GetOwnerName() string {
	return r.OwnerName
}

func (r *Room) GetState() int32 {
	// TODO: Change this to map strings to states
	return 1
}

func (r *Room) GetUsersMax() int32 {
	return int32(r.UsersMax)
}

func (r *Room) GetDescription() string {
	return r.Description
}

func (r *Room) GetScore() int32 {
	return int32(r.Score)
}

func (r *Room) GetCategory() int32 {
	return int32(r.Category)
}

func (r *Room) GetTags() []string {
	return strings.Split(r.Tags, ";")
}

func (r *Room) GetLayout() *RoomLayout {
	return r.layout
}
