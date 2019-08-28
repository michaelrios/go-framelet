package responder_test

import (
	"encoding/json"
	"github.com/michaelrios/go_api/models"
	"net/http"
	"testing"

	"github.com/michaelrios/go_api/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/michaelrios/go_api/api/responder"
	"go.uber.org/zap/zaptest"
)

func TestJsonResponder_RespondWith500(t *testing.T) {
	logger := zaptest.NewLogger(t)
	jsonResponder := responder.NewJsonResponder(logger)

	w := mocks.NewMockWriter()

	jsonResponder.RespondWith500(w)

	assert.Equal(t, http.StatusInternalServerError, w.Assert.Status)

	errors := responder.ErrorResponse{}
	err := json.Unmarshal(w.Assert.Bytes, &errors)
	assert.Nil(t, err)
	assert.Equal(t, "internal server error", errors.Errors[0])
}

func TestJsonResponder_RespondWith400(t *testing.T) {
	logger := zaptest.NewLogger(t)
	jsonResponder := responder.NewJsonResponder(logger)
	w := mocks.NewMockWriter()

	jsonResponder.RespondWith400(w)

	assert.Equal(t, http.StatusBadRequest, w.Assert.Status)

	errors := responder.ErrorResponse{}
	err := json.Unmarshal(w.Assert.Bytes, &errors)
	assert.Nil(t, err)
	assert.Equal(t, "bad request", errors.Errors[0])
}

func TestJsonResponder_RespondWithData(t *testing.T) {
	logger := zaptest.NewLogger(t)
	jsonResponder := responder.NewJsonResponder(logger)
	w := mocks.NewMockWriter()

	jsonResponder.RespondWithData(w, &models.User{UserID: models.UserID("1")})

	assert.Equal(t, http.StatusOK, w.Assert.Status)
	user := models.User{}
	err := json.Unmarshal(w.Assert.Bytes, &user)
	assert.Nil(t, err)
	assert.Equal(t, models.UserID("1"), user.UserID)
}
