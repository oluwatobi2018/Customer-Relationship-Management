package main

import "fmt"

func main() {
	db := NewStorage()
	createUserTable(db.DB)

	fmt.Printf("%#v\n", db)
}
