package middleware_test

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type MockHandler struct {
	Body     []byte
	Status   int
	Response http.ResponseWriter
	Request  *http.Request
	Params   httprouter.Params
}

func (h *MockHandler) Handle(rw http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	h.Response = rw
	h.Request = r
	h.Params = ps

	rw.WriteHeader(h.Status)
	rw.Write(h.Body)
}

type MockResponseWriter struct {
	MockHeader map[string][]string
	MockWrite  struct {
		RecievedBytes []byte
		Output        struct {
			Int   int
			Error error
		}
	}
	ReceivedHeader int
}

func (rp *MockResponseWriter) Header() http.Header {
	return rp.MockHeader
}

func (rp *MockResponseWriter) Write(bytes []byte) (int, error) {
	rp.MockWrite.RecievedBytes = bytes
	return rp.MockWrite.Output.Int, rp.MockWrite.Output.Error
}

func (rp *MockResponseWriter) WriteHeader(statusCode int) {
	rp.ReceivedHeader = statusCode
}
