package middleware

import (
	"go.uber.org/zap"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Recover will recover from a panic in the controller, should be last middleware called
func (m *Middleware) Recover(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer func() {
			if r := recover(); r != nil {
				m.Logger.With(zap.Reflect("recover", r)).DPanic("recover from catastrophic api failure")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("my bad")) // todo make a default error output

				// lookatme: maybe do something here to alert you of this error, cuz yo shit is broke
			}
		}()

		next(w, r, p)
	}
}
