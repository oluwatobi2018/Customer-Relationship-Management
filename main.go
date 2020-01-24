package main

import "fmt"

func main() {
	db := NewStorage()

	fmt.Printf("%#v\n", db)
}
