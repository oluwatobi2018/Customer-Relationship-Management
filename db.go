package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
)

// Storage ...
type Storage struct {
	DB *sql.DB
}

type connData struct {
	host    string
	port    string
	user    string
	pass    string
	dbname  string
	sslmode string
}

// NewStorage ...
func NewStorage() Storage {
	store := Storage{}
	store.connect()
	err := store.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}

	return store
}

func (s *Storage) connect() {
	conn := getENV()
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", conn.host, conn.port, conn.user, conn.pass, conn.dbname, conn.sslmode)
	db, err := sql.Open("postgress", psqlInfo)
	if err != nil {
		log.Fatal(err)
	}
	s.DB = db
}

func getENV() connData {
	conn := connData{}
	var ok bool

	conn.host, ok = os.LookupEnv("DB_HOST")
	if !ok {
		log.Fatal("undefuned `DB_HOST`")
	}
	conn.port, ok = os.LookupEnv("DB_PORT")
	if !ok {
		log.Fatal("undefuned `DB_PORT`")
	}
	conn.user, ok = os.LookupEnv("DB_USER")
	if !ok {
		log.Fatal("undefuned `DB_USER`")
	}
	conn.pass, ok = os.LookupEnv("DB_PASSWORD")
	if !ok {
		log.Fatal("undefuned `DB_PASSWORD`")
	}
	conn.dbname, ok = os.LookupEnv("DBNAME")
	if !ok {
		log.Fatal("undefuned `DB_NAME")
	}
	conn.sslmode, ok = os.LookupEnv("DB_SSLMODE")
	if !ok {
		log.Fatal("undefuned `DB_SSLMODE`")
	}

	return conn
}
