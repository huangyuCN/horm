package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, _ := sql.Open("sqlite3", "test")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("drop table if exists User;")
	_, _ = db.Exec("create table User(Name text);")
	result, err := db.Exec("insert into User(`Name`) values (?),(?)", "Tom", "Sam")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("select Name FROM User limit 1")
	var name string
	if err := row.Scan(&name); err == nil {
		log.Println("name:", name)
	}
}
