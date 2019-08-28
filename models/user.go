package models

import "encoding/json"

type User struct {
	ID       int    `json:"-"`
	UserID   UserID `json:"user_id"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (u *User) Bytes() ([]byte, error) {
	return json.Marshal(u)
}

type UserID string
