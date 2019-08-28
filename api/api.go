package api

import (
	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go_api/api/middleware"
)

// InitializeRoutes initializes the routes given
func InitializeRoutes(routeList []Route, router *httprouter.Router) {
	for _, route := range routeList {
		handler := route.Handler

		for i := len(route.Middleware) - 1; i >= 0; i-- {
			handler = route.Middleware[i](handler)
		}

		getHandlerFunc(route, router)(route.Path, handler)
	}
}

// Just an easy way to talk about your methods, take'em or leave'em
const (
	GET uint8 = iota
	HEAD
	OPTIONS
	POST
	PUT
	PATCH
	DELETE
)

func getHandlerFunc(route Route, router *httprouter.Router) func(string, httprouter.Handle) {
	var handleFunc func(string, httprouter.Handle)
	switch route.Func {
	case GET:
		handleFunc = router.GET
	case HEAD:
		handleFunc = router.HEAD
	case OPTIONS:
		handleFunc = router.OPTIONS
	case POST:
		handleFunc = router.POST
	case PUT:
		handleFunc = router.PUT
	case PATCH:
		handleFunc = router.PATCH
	case DELETE:
		handleFunc = router.DELETE
	default:
		panic("invalid http method")
	}

	return handleFunc
}

type Route struct {
	Func       uint8
	Path       string
	Handler    httprouter.Handle
	Middleware []middleware.Func
}
