package middleware_test

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go_api/api/middleware"
	"go.uber.org/zap/zaptest"
)

func TestMiddleware_AccessLog(t *testing.T) {
	logger := zaptest.NewLogger(t)

	m := middleware.NewMiddleware(logger)

	handler := MockHandler{}

	middlewareFunc := m.AccessLog(handler.Handle)

	responseWriter := &MockResponseWriter{}
	request := &http.Request{
		Method:     http.MethodPost,
		RequestURI: "/some_uri",
		Header: map[string][]string{
			middleware.HeaderXForwardedFor: {"1.1.1.1"},
		},
	}
	params := httprouter.Params{
		httprouter.Param{
			Key:   "someKey",
			Value: "someValue",
		},
	}

	middlewareFunc(responseWriter, request, params)

	assert.Len(t, handler.Params, 1)
	assert.Equal(t, "someKey", handler.Params[0].Key)
	assert.Equal(t, "someValue", handler.Params[0].Value)

	assert.Equal(t, request.RequestURI, handler.Request.RequestURI)

	assert.Len(t, responseWriter.MockHeader, 0)
	assert.Len(t, responseWriter.MockWrite.RecievedBytes, 0)
}
