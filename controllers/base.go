package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go-framelet/dependencies"
	"go.uber.org/zap"
)

func NewSimpleController(deps *dependencies.Core) *SimpleController {
	return &SimpleController{Core: deps}
}

type SimpleController struct {
	*dependencies.Core
}

// Home maybe use this path to return your documentation? ¯\_(ツ)_/¯
func (c *SimpleController) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	logger := c.Logger.With(zap.String("controller", "home"))
	logger.Debug("start")
	defer logger.Debug("done")

	c.RespondWithString(w, "Welcome to another Go template!")
}

// Heartbeat just lets you know if the app is running, should be super simple and fast
func (c *SimpleController) Heartbeat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	logger := c.Logger.With(zap.String("controller", "heartbeat"))
	logger.Debug("start")
	defer logger.Debug("done")

	c.RespondWithString(w, "Yup, still alive.")
}
