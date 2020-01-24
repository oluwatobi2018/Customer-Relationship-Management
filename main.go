package main

import "fmt"

func main() {
	db := NewStorage()
	createUserTable(db.DB, "users")
	createUserTable(db.DB, "address")

	fmt.Printf("%#v\n", db)
}
