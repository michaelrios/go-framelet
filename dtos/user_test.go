package dtos_test

import (
	"testing"

	"github.com/michaelrios/go-framelet/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	user := dtos.NewUser("a@b.c", "abc123")
	assert.Equal(t, "a@b.c", user.Email)
	assert.Equal(t, "abc123", user.Password)
	assert.Equal(t, dtos.UserID(""), user.UserID)
}
