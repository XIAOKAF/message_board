package dao

import "database/sql"

var DB *sql.DB

func InitDB() {
	db, err := sql.Open("mysql", "root:045226@tcp(localhost:3306)/message_board?charset=utf8mb4&parseTime=True")
	if err != nil {
		panic(err)
	}
	DB = db
}
