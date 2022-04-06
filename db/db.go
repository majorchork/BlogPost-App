package db

import (
	"database/sql"
	"log"
)

var DBClient *sql.DB

func OpenDb() {
	db, err := sql.Open("mysql", "root:Login1234@tcp(localhost:3306)/BLOG?parseTime=true")
	if err != nil {
		log.Fatalf(" error connecting to database")
	}
	err = db.Ping()
	if err != nil {
		log.Fatalf(" error pinging database")
	}
	DBClient = db
}
