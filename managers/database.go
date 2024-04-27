package managers

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/alexmgriffiths/habbo-go/utils"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type DatabaseManager struct {
	connection *sql.DB
}

func NewDatabaseManager() *DatabaseManager {

	err := godotenv.Load(".env")
	if err != nil {
		panic(err.Error())
	}
	var dbHost string = os.Getenv("db_host")
	var dbPort string = os.Getenv("db_port")
	var dbUser string = os.Getenv("db_user")
	var dbPass string = os.Getenv("db_pass")
	var dbName string = os.Getenv("db_name")

	var connectionStr string = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	fmt.Println(connectionStr)

	db, err := sql.Open("mysql", connectionStr)

	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	logger := utils.NewLogger()
	logger.Success("Connected to database!")

	return &DatabaseManager{
		connection: db,
	}
}

func (manager *DatabaseManager) GetConnection() *sql.DB {
	return manager.connection
}
