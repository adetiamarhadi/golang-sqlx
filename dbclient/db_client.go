package dbclient

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var DBClient *sql.DB

func InitialiseDBConnection() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}