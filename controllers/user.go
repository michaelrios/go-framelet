package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go-framelet/models"
	"go.uber.org/zap"

	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/domains"
)

func NewUserController(deps *dependencies.Dependencies) *UserController {
	return &UserController{
		Core:       deps.Core,
		UserDomain: domains.NewUserDomain(deps.DB),
	}
}

type UserController struct {
	*dependencies.Core
	*domains.UserDomain
}

func (c *UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	logger := c.Logger.With(zap.String("controller", "GetUser"))
	logger.Debug("start")
	defer logger.Debug("done")

	requestingUser := r.Context().Value("user").(*models.RequestingUser)
	if err := validateUserIsSelf(requestingUser, p.ByName("userid")); err != nil {
		c.RespondWith400(w)
		return
	}

	user, _ := c.UserDomain.GetUser(requestingUser.UserID)

	c.RespondWithData(w, user)
}

func validateUserIsSelf(requestingUser *models.RequestingUser, requestedUser string) error {
	if requestingUser.IsEmpty() {
		return fmt.Errorf("")
	}

	return nil
}
