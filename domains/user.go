package domains

import (
	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/dtos"
	"github.com/michaelrios/go-framelet/repositories"
	"golang.org/x/xerrors"
)

var UserNotFound = xerrors.New("user not found")

func NewUserDomain(db *dependencies.DB) *UserDomain {
	return &UserDomain{userRepository: repositories.NewUserRepository(db)}
}

type UserDomain struct {
	userRepository *repositories.UserRepository
}

func (d *UserDomain) GetUser(userId dtos.UserID) (*dtos.User, error) {
	user, err := d.userRepository.GetUser(userId)
	if err == repositories.ErrEntityNotFound {
		err = xerrors.Errorf("user not found: %w", err)
		return nil, err
	}

	return user, nil
}

func (d *UserDomain) CreateUser(user *dtos.User) (*dtos.User, error) {
	user, err := d.userRepository.CreateUser(user)
	if err == repositories.ErrEntityNotFound {
		err = xerrors.Errorf("user not found: %w", err)
		return nil, err
	}

	// produce new user

	return user, nil
}

func (d *UserDomain) UpdateUser(user *dtos.User) (*dtos.User, error) {
	user, err := d.userRepository.UpdateUser(user)
	if err == repositories.ErrEntityNotFound {
		err = xerrors.Errorf("user not found: %w", err)
		return nil, err
	}

	// produce updated user

	return user, nil
}

func (d *UserDomain) DeleteUser(userID dtos.UserID) error {
	err := d.userRepository.DeleteUser(userID)
	if err == repositories.ErrEntityNotFound {
		err = xerrors.Errorf("user not found: %w", err)
		return err
	}

	// produce deleted user

	return nil
}
