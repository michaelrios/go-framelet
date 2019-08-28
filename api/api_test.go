package api_test

import (
	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go-framelet/api"
	"github.com/michaelrios/go-framelet/api/middleware"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap/zaptest"
	"net/http"
	"testing"
)

func TestInitializeRoutes(t *testing.T) {
	logger := zaptest.NewLogger(t)
	testRouter :=  httprouter.New()

	m := middleware.NewMiddleware(logger)

	routes := getTestRoutes(m)

	api.InitializeRoutes(routes, testRouter)

	_, _, isSet := testRouter.Lookup("GET", "")
	assert.True(t, isSet)
}

func MockHandler(http.ResponseWriter, *http.Request, httprouter.Params) {}

func getTestRoutes(m *middleware.Middleware) []api.Route {
	return []api.Route{
		{
			Func: api.GET,
			Path: "/",
			Handler: MockHandler,
			Middleware: []middleware.Func{m.AccessLog},
		},
		{
			Func: api.HEAD,
			Path: "/",
			Handler: MockHandler,
			Middleware: []middleware.Func{m.AccessLog},
		},
		{
			Func: api.OPTIONS,
			Path: "/",
			Handler: MockHandler,
			Middleware: []middleware.Func{m.AccessLog},
		},
		{
			Func: api.POST,
			Path: "/",
			Handler: MockHandler,
			Middleware: []middleware.Func{m.AccessLog},
		},
		{
			Func: api.PUT,
			Path: "/",
			Handler: MockHandler,
			Middleware: []middleware.Func{m.AccessLog},
		},
		{
			Func: api.PATCH,
			Path: "/",
			Handler: MockHandler,
			Middleware: []middleware.Func{m.AccessLog},
		},
		{
			Func: api.DELETE,
			Path: "/",
			Handler: MockHandler,
			Middleware: []middleware.Func{m.AccessLog},
		},
	}
}