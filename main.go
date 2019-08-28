package main

import (
	"flag"
	"net/http"

	"github.com/michaelrios/go_api/dependencies"

	"github.com/pkg/profile"

	"github.com/julienschmidt/httprouter"
	"github.com/michaelrios/go_api/api"
	"go.uber.org/zap"
)

var shouldProfile = flag.Bool("p", false, "profile application")

func main() {
	flag.Parse()
	if *shouldProfile { // lookatme: this is just for easier profiling, feel free to remove
		defer profile.Start().Stop()
	}

	// Config stuff
	config, err := ParseConfigs()
	if err != nil {
		panic("failed to parse configs" + err.Error())
	}

	// logger stuff
	logger, err := zap.NewDevelopment()
	if config.Logger.Structured {
		logger, err = zap.NewProduction()
	}
	if err != nil {
		panic("logger failed to start" + err.Error())
	}

	// todo: add DB examples like Mongo and Redis

	// make repository things

	// done making dependency things
	logger.Info("APP dependencies are initialized")

	router := httprouter.New()

	deps := &dependencies.Dependencies{
		Core: dependencies.NewCore(logger),
		DB:   &dependencies.DB{},
	}

	api.InitializeRoutes(BuildRoutes(deps), router)

	logger.Panic("router stopped",
		zap.Error(http.ListenAndServe(config.Server.Host+config.Server.Port, router)),
	)
}
