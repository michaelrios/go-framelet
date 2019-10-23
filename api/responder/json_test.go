package responder_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/michaelrios/go-framelet/controllers/viewmodels"
	"github.com/michaelrios/go-framelet/dtos"

	"github.com/michaelrios/go-framelet/mocks"
	"github.com/stretchr/testify/assert"

	"github.com/michaelrios/go-framelet/api/responder"
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

	jsonResponder.RespondWithData(w, &viewmodels.ResponseUser{UserID: "1"})

	assert.Equal(t, http.StatusOK, w.Assert.Status)
	user := viewmodels.ResponseUser{}
	err := json.Unmarshal(w.Assert.Bytes, &user)
	assert.Nil(t, err)
	assert.Equal(t, dtos.UserID("1"), user.UserID)
}
