package main

import "database/sql"

func conn() (*sql.DB, error) {
	db, err := sql.Open("mysql")
}
