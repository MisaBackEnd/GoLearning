package main

import (
	"encoding/json"
	"net/http"
)

// es una funcion de tipo handler
type Middleware func(http.HandlerFunc) http.HandlerFunc

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func (u *User) userToJson() ([]byte, error) {
	return json.Marshal(u)
}
