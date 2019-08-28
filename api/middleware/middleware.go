package middleware

import (
	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go-framelet/dependencies"
	"go.uber.org/zap"
)

type Func func(handle httprouter.Handle) httprouter.Handle

type Middleware struct {
	*dependencies.Core
}

func NewMiddleware(logger *zap.Logger) *Middleware {
	return &Middleware{
		Core: &dependencies.Core{Logger: logger},
	}
}
