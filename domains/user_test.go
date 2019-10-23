package domains_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/domains"
)

func TestNewUserDomain(t *testing.T) {
	userDomain := domains.NewUserDomain(&dependencies.DB{})
	assert.IsType(t, userDomain, new(domains.UserDomain))
}
