package dtos

type UserID string

type User struct {
	UserID   UserID
	Email    string
	Password string
}

func NewUser(email string, password string) User {
	return User{Email: email, Password: password}
}
