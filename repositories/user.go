package repositories

import (
	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/models"
)

func NewUserRepository(db *dependencies.DB) *UserRepository {
	return &UserRepository{db: db}
}

type UserRepository struct {
	db *dependencies.DB
}

func (ur *UserRepository) GetUser(id models.UserID) (*models.User, error) {
	return &models.User{UserID: id, Email: "hi@email.com"}, nil
}
