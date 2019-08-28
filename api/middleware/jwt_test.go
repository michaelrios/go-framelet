package middleware_test

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go-framelet/api/middleware"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest"
	"net/http"
	"testing"
	"time"
)

func TestMiddleware_JWT_Success(t *testing.T) {
	logger := zaptest.NewLogger(t)

	m := middleware.NewMiddleware(logger)
	handler := MockHandler{}
	middlewareFunc := m.JWT(handler.Handle)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, middleware.Claims{
		UserId: "123",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			Subject: "subject",
			IssuedAt: time.Now().Unix(),
		},
	})
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		logger.With(zap.Error(err)).Error("jwt signing error")
	}

	responseWriter := &MockResponseWriter{}
	request := &http.Request{
		Header: map[string][]string{
			"Authorization": {"Bearer "+tokenString},
		},
	}

	middlewareFunc(responseWriter, request, httprouter.Params{})

	assert.Len(t, responseWriter.MockWrite.RecievedBytes, 0)
	assert.Equal(t, responseWriter.ReceivedHeader, 0)
}

func TestMiddleware_JWT_Expired(t *testing.T) {
	logger := zaptest.NewLogger(t)

	logger.WithOptions()

	m := middleware.NewMiddleware(logger)

	handler := MockHandler{}

	middlewareFunc := m.JWT(handler.Handle)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, middleware.Claims{
		UserId: "123",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(-time.Minute).Unix(),
			Subject: "subject",
			IssuedAt: time.Now().Unix(),
		},
	})
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		logger.With(zap.Error(err)).Error("jwt signing error")
	}

	responseWriter := &MockResponseWriter{}
	request := &http.Request{
		Header: map[string][]string{
			"Authorization": {"Bearer "+tokenString},
		},
	}

	middlewareFunc(responseWriter, request, httprouter.Params{})

	assert.Len(t, responseWriter.MockWrite.RecievedBytes, 4)
	assert.Equal(t, responseWriter.ReceivedHeader, 400)
}

func TestMiddleware_JWT_MissingUser(t *testing.T) {
	logger := zaptest.NewLogger(t)

	logger.WithOptions()

	m := middleware.NewMiddleware(logger)

	handler := MockHandler{}

	middlewareFunc := m.JWT(handler.Handle)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, middleware.Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute).Unix(),
			Subject: "subject",
			IssuedAt: time.Now().Unix(),
		},
	})
	tokenString, err := token.SignedString(middleware.JwtKey)
	if err != nil {
		logger.With(zap.Error(err)).Error("jwt signing error")
	}

	responseWriter := &MockResponseWriter{}
	request := &http.Request{
		Header: map[string][]string{
			"Authorization": {"Bearer "+tokenString},
		},
	}

	middlewareFunc(responseWriter, request, httprouter.Params{})

	assert.Len(t, responseWriter.MockWrite.RecievedBytes, 4)
	assert.Equal(t, responseWriter.ReceivedHeader, 400)
}

func TestMiddleware_JWT_ParseFailure(t *testing.T) {
	logger := zaptest.NewLogger(t)

	logger.WithOptions()

	m := middleware.NewMiddleware(logger)

	handler := MockHandler{}

	middlewareFunc := m.JWT(handler.Handle)

	responseWriter := &MockResponseWriter{}
	request := &http.Request{
		Header: map[string][]string{
			"Authorization": {""},
		},
	}

	params := httprouter.Params{}

	middlewareFunc(responseWriter, request, params)

	assert.Len(t, responseWriter.MockWrite.RecievedBytes, 4)
	assert.Equal(t, responseWriter.ReceivedHeader, 400)
}

func TestMiddleware_JWT_MissingAuth(t *testing.T) {
	logger := zaptest.NewLogger(t)

	logger.WithOptions()

	m := middleware.NewMiddleware(logger)

	handler := MockHandler{}

	middlewareFunc := m.JWT(handler.Handle)

	responseWriter := &MockResponseWriter{}
	request := &http.Request{
		Header: map[string][]string{},
	}

	params := httprouter.Params{}

	middlewareFunc(responseWriter, request, params)

	assert.Len(t, responseWriter.MockWrite.RecievedBytes, 4)
	assert.Equal(t, responseWriter.ReceivedHeader, 400)
}
