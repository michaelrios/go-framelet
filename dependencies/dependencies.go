package dependencies

import (
	"github.com/michaelrios/go-framelet/api/responder"
	"go.uber.org/zap"
)

type Dependencies struct {
	*Core
	DB *DB
}

type Core struct {
	Logger *zap.Logger
	*responder.JsonResponder
}

func NewCore(logger *zap.Logger) *Core {
	return &Core{Logger: logger, JsonResponder: responder.NewJsonResponder(logger)}
}
