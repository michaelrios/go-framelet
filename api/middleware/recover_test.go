package middleware_test

import (
	"net/http"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go-framelet/api/middleware"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
)

func TestMiddleware_Recover(t *testing.T) {
	logger := zaptest.NewLogger(t)

	m := middleware.NewMiddleware(logger)

	handler := MockHandler{}

	middlewareFunc := m.Recover(handler.Handle)

	responseWriter := &MockResponseWriter{}
	request := &http.Request{}
	params := httprouter.Params{}

	middlewareFunc(responseWriter, request, params)

	assert.Len(t, handler.Params, 0)
	assert.Len(t, responseWriter.MockHeader, 0)
	assert.Len(t, responseWriter.MockWrite.RecievedBytes, 0)
}
