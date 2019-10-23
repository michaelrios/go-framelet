package models

import "github.com/michaelrios/go-framelet/dtos"

type User struct {
	ID       uint
	UserID   dtos.UserID
	Email    string
	Password string
}
