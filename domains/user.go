package domains

import (
	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/models"
	"github.com/michaelrios/go-framelet/repositories"
	"golang.org/x/xerrors"
)

var UserNotFound = xerrors.New("user not found")

func NewUserDomain(db *dependencies.DB) *UserDomain {
	return &UserDomain{repository: repositories.NewUserRepository(db)}
}

type UserDomain struct {
	repository *repositories.UserRepository
}

func (d *UserDomain) GetUser(userId models.UserID) (*models.User, error) {
	user, err := d.repository.GetUser(userId)
	if err == repositories.ErrEntityNotFound {
		err = xerrors.Errorf("user not found: %w", err)
		return nil, err
	} else {
		return nil, err
	}

	return user, nil
}
