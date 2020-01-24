package main

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
