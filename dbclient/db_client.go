package dbclient

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var DBClient *sqlx.DB

func InitialiseDBConnection() {
	db, err := sqlx.Open("mysql", "root:root@tcp(localhost:3306)/test_db?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	err = db.Ping()
	if err != nil {
		panic(err.Error())
	}

	DBClient = db
}