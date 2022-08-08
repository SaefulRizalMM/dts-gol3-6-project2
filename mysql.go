package main

import (
	"database/sql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/dts_gol3_6_project2")

	if err != nil {
		log.Fatal(err)
	}

	return db
}
