package viewmodels_test

import (
	"testing"

	"github.com/michaelrios/go-framelet/dtos"

	"github.com/michaelrios/go-framelet/controllers/viewmodels"
	"github.com/stretchr/testify/assert"
)

func TestRequestUserToDTO(t *testing.T) {
	dto := viewmodels.RequestUserToDTO(&viewmodels.RequestUser{})
	assert.Equal(t, dtos.UserID(""), dto.UserID)
	assert.Equal(t, "", dto.Email)

	dto = viewmodels.RequestUserToDTO(&viewmodels.RequestUser{UserID: "1", Email: "a@b.c"})
	assert.Equal(t, dtos.UserID("1"), dto.UserID)
	assert.Equal(t, "a@b.c", dto.Email)
}

func TestResponseUserFromDTO(t *testing.T) {
	responseUser := viewmodels.ResponseUserFromDTO(dtos.User{})
	assert.Equal(t, dtos.UserID(""), responseUser.UserID)
	assert.Equal(t, "", responseUser.Email)

	responseUser = viewmodels.ResponseUserFromDTO(dtos.User{UserID: "1", Email: "a@b.c", Password: "abc123"})
	assert.Equal(t, dtos.UserID("1"), responseUser.UserID)
	assert.Equal(t, "a@b.c", responseUser.Email)
}
