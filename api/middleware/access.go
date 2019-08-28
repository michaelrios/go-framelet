package middleware

import (
	"net"
	"net/http"
	"strings"
	"time"

	"go.uber.org/zap/zapcore"

	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
)

const (
	HeaderXForwardedFor = "X-Forwarded-For"
	HeaderXRealIP       = "X-Real-IP"
)

// AccessLog logs all http requests and should be the first middleware ran
func (m *Middleware) AccessLog(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		m.Logger.Check(zapcore.DebugLevel, "middleware: access start").Write()
		defer m.Logger.Check(zapcore.DebugLevel, "middleware: access done").Write()
		if r.RequestURI == "/favicon.ico" { // lookatme: this annoyed me to see it when called by the browser, up to you if you want to keep it
			next(w, r, p)
			return
		}
		rw := &ResponseWriter{
			ResponseWriter: w,
		}
		start := time.Now()

		defer func() {
			m.Logger.With(
				zap.String("bytes_in", r.Header.Get("Content-Length")),
				zap.Int("bytes_out", rw.BytesOut),
				zap.String("duration", time.Now().Sub(start).String()),
				zap.String("remote_ip", realIP(r)),
				zap.String("host", r.Host),
				zap.String("uri", r.RequestURI),
				zap.String("method", r.Method),
				zap.Int("status", rw.Status)).
				Info("access_log")
		}()

		next(rw, r, p)
	}
}

// ResponseWriter hijacks the default ResponseWriter to track the status and bytesOut
type ResponseWriter struct {
	http.ResponseWriter
	Status   int
	BytesOut int
}

// Write writes bytes to the response out
func (rw *ResponseWriter) Write(bytes []byte) (int, error) {
	bytesOut, err := rw.ResponseWriter.Write(bytes)
	rw.BytesOut += bytesOut
	return bytesOut, err
}

// WriteHeader sets the status code to the response
func (rw *ResponseWriter) WriteHeader(statusCode int) {
	rw.Status = statusCode
	rw.ResponseWriter.WriteHeader(statusCode)
}

func realIP(r *http.Request) string {
	if ip := r.Header.Get(HeaderXForwardedFor); ip != "" {
		return strings.Split(ip, ", ")[0]
	}
	if ip := r.Header.Get(HeaderXRealIP); ip != "" {
		return ip
	}
	ra, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ra
}
