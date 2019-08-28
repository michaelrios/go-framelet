package domains

import (
	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/models"
	"github.com/michaelrios/go-framelet/repositories"
)

func NewUserDomain(db *dependencies.DB) *UserDomain {
	return &UserDomain{repository: repositories.NewUserRepository(db)}
}

type UserDomain struct {
	repository *repositories.UserRepository
}

func (d *UserDomain) GetUser(userId models.UserID) (*models.User, error) {
	return d.repository.GetUser(userId)
}
