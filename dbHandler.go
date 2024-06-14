package main

import "database/sql"

type DBMaker struct {
	db *sql.DB
}

func DBConnection() (schema *DBMaker, err error) {
	connectionString := "user=postgres dbname=postgres password=user ssl=disable"
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &DBMaker{
		db: db,
	}, nil
}
