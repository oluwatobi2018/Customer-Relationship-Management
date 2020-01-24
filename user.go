package main

import (
	"database/sql"
	"log"
)

// User ...
type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Profile ...
type Profile struct {
	Position string `json:"position"`
	Phone    string `json:"phone"`
	Address  Addr   `json:"address"`
}

// Addr ...
type Addr struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	Country string `json:"country"`
}

func createUserTable(db *sql.DB, name string) {
	switch name {
	case "users":
		createUsers(db)
	case "address":
		createAddress(db)
	}
}

func createUsers(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name character varying(256),
		email character varying(256)
		password character varying(256)
		password character varying(256)
		password character varying(256)
	);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS users_pkey ON users(id int4_ops);")
	if err != nil {
		log.Fatal(err)
	}

}

func createAddress(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE IF NOT EXISTS address (
		id SERIAL PRIMARY KEY,
		street character varying(256)
		city character varying(256)
		country character varying(256)
	);`)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec("CREATE UNIQUE INDEX IF NOT EXISTS users_pkey ON users(id int4_ops);")
	if err != nil {
		log.Fatal(err)
	}

}
