package main

import (
	"github.com/michaelrios/go-framelet/api"
	"github.com/michaelrios/go-framelet/api/middleware"
	"github.com/michaelrios/go-framelet/controllers"
	"github.com/michaelrios/go-framelet/dependencies"
)

// BuildRoutes returns all of your routes
func BuildRoutes(deps *dependencies.Dependencies) []api.Route { // lookatme: this should list your http routes
	m := middleware.Middleware{Core: deps.Core}

	simple := controllers.NewSimpleController(deps.Core)

	user := controllers.NewUserController(deps)

	return []api.Route{
		{
			Func:       api.GET,
			Path:       "/",
			Handler:    simple.Home,
			Middleware: []middleware.Func{m.AccessLog, m.Recover},
		},
		{
			Func:       api.GET,
			Path:       "/heartbeat",
			Handler:    simple.Heartbeat,
			Middleware: []middleware.Func{m.AccessLog, m.Recover},
		},
		{
			Func:       api.GET,
			Path:       "/users", // todo what should the path be here to differentiate between GetUser and GetUsers
			Handler:    user.GetUser,
			Middleware: []middleware.Func{m.AccessLog, m.JWT, m.Recover},
		},
	}
}
