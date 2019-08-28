package responder

import (
	"encoding/json"
	"net/http"

	"github.com/michaelrios/go-framelet/models"
	"go.uber.org/zap"
)

func NewJsonResponder(logger *zap.Logger) *JsonResponder {
	return &JsonResponder{logger: logger}
}

type JsonResponder struct {
	logger *zap.Logger
}

func (r *JsonResponder) RespondWithData(w http.ResponseWriter, viewable models.Viewable) {
	bytes, err := viewable.Bytes()
	if err != nil {
		r.RespondWith500(w)
		return
	}

	r.respondWithJson(w, http.StatusOK, bytes)
}

func (r *JsonResponder) RespondWithString(w http.ResponseWriter, msg string) {
	r.respondWithText(w, http.StatusOK, []byte(msg))
}

func (r *JsonResponder) RespondWith400(w http.ResponseWriter, errors ...string) {
	if len(errors) == 0 {
		errors = append(errors, "bad request")
	}

	errorResponse := ErrorResponse{Errors: errors}
	bytes, err := json.Marshal(errorResponse)
	if err != nil {
		r.logger.With(zap.Error(err)).DPanic("failed to marshal data")
		r.RespondWith500(w)
		return
	}

	r.respondWithJson(w, http.StatusBadRequest, bytes)
}

func (r *JsonResponder) RespondWith500(w http.ResponseWriter) {
	r.respondWithJson(w, http.StatusInternalServerError, []byte("{\"errors\":[\"internal server error\"]}"))
}

type ErrorResponse struct {
	Errors []string `json:"errors"`
}

func (r *ErrorResponse) Bytes() ([]byte, error) {
	return json.Marshal(r)
}

func (r *JsonResponder) respondWithJson(w http.ResponseWriter, status int, bytes []byte) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")

	if _, err := w.Write(bytes); err != nil {
		r.logger.With(zap.Error(err)).DPanic("failed to write json response")
	}
}

func (r *JsonResponder) respondWithText(w http.ResponseWriter, status int, msg []byte) {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "text/plain")

	if _, err := w.Write(msg); err != nil {
		r.logger.With(zap.Error(err)).DPanic("failed to write json response")
	}
}
