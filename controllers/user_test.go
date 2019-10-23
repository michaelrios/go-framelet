package controllers_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/michaelrios/go-framelet/api/middleware"

	"go.uber.org/zap/zaptest"

	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go-framelet/controllers"
	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/mocks"
	"github.com/stretchr/testify/assert"
)

func TestUserController_GetUser(t *testing.T) {
	logger := zaptest.NewLogger(t)
	userController := controllers.NewUserController(&dependencies.Dependencies{
		Core: &dependencies.Core{
			Logger: logger,
		},
	})

	request := httptest.NewRequest(http.MethodGet, "/users", strings.NewReader(""))
	ctx := context.WithValue(context.Background(), "user", &middleware.AuthenticatedUser{UserID: "1"})

	responseWriter := mocks.NewMockWriter()
	params := httprouter.Params{}
	params = append(params, httprouter.Param{Key: "userid", Value: "1"})

	userController.GetUser(responseWriter, request.WithContext(ctx), params)

	assert.Equal(t, 200, responseWriter.Assert.Status)
	assert.Equal(t, `{"user_id":"1","email":"a@b.c"}`, string(responseWriter.Assert.Bytes))
}

func TestUserController_CreateUser(t *testing.T) {
	logger := zaptest.NewLogger(t)
	userController := controllers.NewUserController(&dependencies.Dependencies{
		Core: &dependencies.Core{
			Logger: logger,
		},
	})

	request := httptest.NewRequest(http.MethodPost, "/users", strings.NewReader(`{"email":"a@b.c","password":"Abc123!"}`))

	responseWriter := mocks.NewMockWriter()
	params := httprouter.Params{}

	userController.CreateUser(responseWriter, request, params)

	assert.Equal(t, 200, responseWriter.Assert.Status)
	assert.Equal(t, `{"user_id":"1","email":"a@b.c"}`, string(responseWriter.Assert.Bytes))
}

func TestUserController_UpdateUser(t *testing.T) {
	logger := zaptest.NewLogger(t)
	userController := controllers.NewUserController(&dependencies.Dependencies{
		Core: &dependencies.Core{
			Logger: logger,
		},
	})

	request := httptest.NewRequest(http.MethodPut, "/users", strings.NewReader(`{"email":"a@b.c","password":"Abc123!"}`))

	responseWriter := mocks.NewMockWriter()
	params := httprouter.Params{}

	userController.UpdateUser(responseWriter, request, params)

	assert.Equal(t, 200, responseWriter.Assert.Status)
	assert.Equal(t, `{"user_id":"1","email":"a@b.c"}`, string(responseWriter.Assert.Bytes))
}
