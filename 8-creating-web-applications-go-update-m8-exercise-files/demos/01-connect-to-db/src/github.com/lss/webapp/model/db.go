package model

import "database/sql"

var db *sql.DB

func SetDatabase(database *sql.DB) {
	db = database
}
