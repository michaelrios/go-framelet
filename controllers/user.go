package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/michaelrios/go-framelet/api/middleware"
	"github.com/michaelrios/go-framelet/controllers/viewmodels"
	"github.com/michaelrios/go-framelet/dependencies"
	"github.com/michaelrios/go-framelet/domains"
	"github.com/michaelrios/go-framelet/dtos"
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

	authenticatedUser := r.Context().Value("user").(*middleware.AuthenticatedUser)
	requestedUserID := p.ByName("userid")
	if err := validateUserIsSelf(authenticatedUser, requestedUserID); err != nil {
		logger.With(zap.Error(err)).With(zap.String("requested_user", requestedUserID)).Info("failed validating user")
		c.RespondWith400(w, "invalid user")
		return
	}

	user, err := c.UserDomain.GetUser(dtos.UserID(authenticatedUser.UserID))
	if xerrors.Is(err, domains.UserNotFound) {
		c.RespondWith404(w)
		return
	} else if err != nil {
		c.Logger.With(zap.Error(err)).Error("failed getting user")
		c.RespondWith500(w)
		return
	}

	c.RespondWithData(w, viewmodels.ResponseUserFromDTO(*user))
}

func (c *UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	logger := c.Logger.With(zap.String("controller", "CreateUser"))
	logger.Debug("start")
	defer logger.Debug("done")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.RespondWith500(w)
	}

	requestUser := &viewmodels.RequestUser{}
	if err := json.Unmarshal(bytes, requestUser); err != nil {
		c.RespondWith400(w, "invalid json body")
	}

	createdUser, err := c.UserDomain.CreateUser(viewmodels.RequestUserToDTO(requestUser))
	if xerrors.Is(err, domains.UserNotFound) {
		c.Logger.With(zap.Error(err)).
			With(zap.String("user_id", string(requestUser.UserID))).
			Info("user not found")
		c.RespondWith404(w)
		return
	} else if err != nil {
		c.Logger.With(zap.Error(err)).Error("failed getting user")
		c.RespondWith500(w)
		return
	}

	c.RespondWithData(w, viewmodels.ResponseUserFromDTO(*createdUser))
}

func (c *UserController) UpdateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	logger := c.Logger.With(zap.String("controller", "UpdateUser"))
	logger.Debug("start")
	defer logger.Debug("done")

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		c.RespondWith400(w, "")
	}

	requestUser := &viewmodels.RequestUser{}
	if err := json.Unmarshal(bytes, requestUser); err != nil {
		c.RespondWith400(w, "")
	}

	createdUser, err := c.UserDomain.CreateUser(viewmodels.RequestUserToDTO(requestUser))
	if xerrors.Is(err, domains.UserNotFound) {
		c.RespondWith404(w)
		return
	} else if err != nil {
		c.Logger.With(zap.Error(err)).Error("failed getting user")
		c.RespondWith500(w)
		return
	}

	c.RespondWithData(w, viewmodels.ResponseUserFromDTO(*createdUser))
}

func validateUserIsSelf(requestingUser *middleware.AuthenticatedUser, requestedUser string) error {
	if requestingUser.IsEmpty() {
		return xerrors.Errorf("no requesting user given")
	}

	if requestingUser.UserID != requestedUser {
		return xerrors.Errorf("requesting user is not requested user")
	}

	return nil
}
