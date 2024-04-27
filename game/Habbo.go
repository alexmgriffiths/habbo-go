package game

import (
	"database/sql"
	"strings"

	"github.com/gorilla/websocket"
)

type Habbo struct {
	connection *websocket.Conn
	id         int32
	username   string
	authTicket string
	look       string
	credits    int32
	motto      string
	gender     string
}

func NewHabbo(db *sql.DB, conn *websocket.Conn, authTicket string) (*Habbo, error) {

	row := db.QueryRow("SELECT * FROM users WHERE auth_ticket = ?", authTicket)

	var id int32
	var username string
	var real_name sql.NullString
	var password string
	var mail sql.NullString
	var mail_verified byte
	var account_created sql.NullInt32
	var account_day_of_birth int32
	var last_login int32
	var last_online int32
	var motto string
	var look string
	var gender string
	var rank int8
	var credits int32
	var pixels int32
	var points int32
	var online int8
	var auth_ticket string
	var ip_register sql.NullString
	var ip_current sql.NullString
	var machine_id string
	var home_room int32
	var secret_key sql.NullString
	var pincode sql.NullString
	var extra_rank sql.NullInt16

	err := row.Scan(
		&id,
		&username,
		&real_name,
		&password,
		&mail,
		&mail_verified,
		&account_created,
		&account_day_of_birth,
		&last_login,
		&last_online,
		&motto,
		&look,
		&gender,
		&rank,
		&credits,
		&pixels,
		&points,
		&online,
		&auth_ticket,
		&ip_register,
		&ip_current,
		&machine_id,
		&home_room,
		&secret_key,
		&pincode,
		&extra_rank,
	)

	if err != nil {
		return nil, err
	}

	println(id, username)

	return &Habbo{
		connection: conn,
		id:         id,
		username:   username,
		authTicket: authTicket,
		gender:     gender,
		look:       look,
		credits:    credits,
		motto:      motto,
	}, nil
}

func (h *Habbo) GetConnection() *websocket.Conn {
	return h.connection
}

func (h *Habbo) GetId() int32 {
	return h.id
}

func (h *Habbo) GetUsername() string {
	return h.username
}

func (h *Habbo) GetLook() string {
	return h.look
}

func (h *Habbo) GetGender() string {
	return strings.ToUpper(h.gender)
}

func (h *Habbo) GetMotto() string {
	return h.motto
}
