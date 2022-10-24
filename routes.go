package main


import (
	"database/sql"
	"net/http"
)

// UserHandler ...
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		UserList(w, r)
	case "POST":
		UserAdd(w, r)
	}
}

// UserList ...
func UserList(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sql.DB)
	UserObj{}.RredAll(db)
	w.Write([]byte("ok\n"))
}

// UserAdd ...
func UserAdd(w http.ResponseWriter, r *http.Request) {
	db := r.Context().Value("db").(*sql.DB)
	UserObj{}.Add(db, User{
		Name: "max",
	})
	w.Write([]byte("ok\n"))
}
