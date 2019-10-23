package repositories_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/dtos"
	"github.com/michaelrios/go-framelet/repositories"
)

func TestUserRepository_CreateUser(t *testing.T) {
	db := &dependencies.DB{}

	userRepository := repositories.NewUserRepository(db)

	userDTO, err := userRepository.GetUser(dtos.UserID("1"))
	assert.Nil(t, err)

	assert.Equal(t, dtos.UserID("1"), userDTO.UserID)
}
