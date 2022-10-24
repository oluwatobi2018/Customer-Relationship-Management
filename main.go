package main

import (
	"context"
	"log"
	"net/http"
	"os"
)

// ContextInjector - inject DB
type ContextInjector struct {
	ctx context.Context
	h   http.Handler
}

func (ci *ContextInjector) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ci.h.ServeHTTP(w, r.WithContext(ci.ctx))
}

func main() {
	db := NewStorage()
	createUserTable(db.DB, "users")
	createUserTable(db.DB, "address")

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = "localhost"
		log.Print("using default host `localhost`")
	}
	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = "8080"
		log.Print("using default port `80`")
	}

	// create base Context
	ctx := context.WithValue(context.Background(), "db", db.DB)

	// handlers

	// main
	http.Handle("/", &ContextInjector{ctx, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("ok"))
	})})

	// users handlers
	http.Handle("/user", http.Handler(&ContextInjector{ctx, http.HandlerFunc(UserHandler)}))

	log.Fatal(http.ListenAndServe(host+":"+port, nil))
}

