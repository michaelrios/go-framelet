package repositories

import (
	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/dtos"
)

func NewUserRepository(db *dependencies.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepository struct {
	db *dependencies.DB
}

func (ur *UserRepository) GetUser(id dtos.UserID) (*dtos.User, error) {
	return &dtos.User{UserID: "1", Email: "a@b.c"}, nil
}

func (ur *UserRepository) CreateUser(user *dtos.User) (*dtos.User, error) {
	return &dtos.User{UserID: "1", Email: "a@b.c"}, nil
}

func (ur *UserRepository) UpdateUser(user *dtos.User) (*dtos.User, error) {
	return user, nil
}

func (ur *UserRepository) DeleteUser(id dtos.UserID) error {
	return nil
}
