package rooms

type RoomCategory struct {
	id           int
	minRank      int
	captionSave  string
	caption      string
	canTrade     int
	maxUserCount int
	public       int
	listType     int
	orderNum     int
}

func NewRoomCategory(id int, minRank int, captionSave string, caption string, canTrade int, maxUserCount int, public int, listType int, orderNum int) *RoomCategory {

	return &RoomCategory{
		id:           id,
		minRank:      minRank,
		captionSave:  captionSave,
		caption:      caption,
		canTrade:     canTrade,
		maxUserCount: maxUserCount,
		public:       public,
		listType:     listType,
		orderNum:     orderNum,
	}

}

func (rc *RoomCategory) GetId() int {
	return rc.id
}

func (rc *RoomCategory) GetCaption() string {
	return rc.caption
}
