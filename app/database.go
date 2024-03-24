package app

import (
	"database/sql"
	"time"
	"todo-api/helper"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/todo_api_golang?parseTime=true")

	helper.PanicIfError(err)

	db.SetConnMaxIdleTime(5)
	db.SetMaxOpenConns(10)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
	// migrate -database "mysql://root:root@tcp(localhost:3306)/todo_api_golang" -path db/migrations up
}
