package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/michaelrios/go-framelet/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/julienschmidt/httprouter"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var JwtKey = []byte("your-256-bit-secret")

// JWT parses the given Authorization token in the request
func (m *Middleware) JWT(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		m.Logger.Check(zapcore.DebugLevel, "middleware: jwt start").Write()
		defer m.Logger.Check(zapcore.DebugLevel, "middleware: jwt done").Write()

		claims := &Claims{}
		jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		token := strings.TrimSpace(strings.TrimLeft(r.Header.Get("Authorization"), "Bearer"))

		if _, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
			return JwtKey, nil
		}); err != nil {
			m.Logger.With(zap.Error(err)).Debug("failed to parse jwt")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("oops")) // todo make a default error output
			return
		}

		if claims.UserId == "" {
			m.Logger.Debug("jwt: user not set")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("oops")) // todo make a default error output
			return
		}

		m.Logger.With(zap.Reflect("user_id", claims.UserId)).Info("jwt: successfully parsed")

		ctx := context.WithValue(r.Context(), "user", claims.RequestingUser())

		next(w, r.WithContext(ctx), p)
	}
}

// Claims from JWT contents
type Claims struct {
	UserId      models.UserID `json:"user_id"`
	Permissions []string      `json:"permissions"`
	jwt.StandardClaims
}

func (c *Claims) RequestingUser() *models.RequestingUser {
	permissions := make(map[string]bool)
	for _, v := range c.Permissions {
		permissions[v] = true
	}

	return &models.RequestingUser{
		UserID:      c.UserId,
		Permissions: permissions,
	}
}
