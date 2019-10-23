package viewmodels

import (
	"encoding/json"

	"github.com/michaelrios/go-framelet/dtos"
)

type RequestUser struct {
	UserID   dtos.UserID `json:"user_id"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
}

func RequestUserToDTO(user *RequestUser) *dtos.User {
	return &dtos.User{
		UserID: user.UserID,
		Email:  user.Email,
	}
}

type ResponseUser struct {
	UserID dtos.UserID `json:"user_id"`
	Email  string      `json:"email"`
}

func (u *ResponseUser) Bytes() ([]byte, error) {
	return json.Marshal(u)
}

func ResponseUserFromDTO(user dtos.User) *ResponseUser {
	return &ResponseUser{
		UserID: user.UserID,
		Email:  user.Email,
	}
}
