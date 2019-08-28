package main

import "github.com/caarlos0/env"

// Config describes the environment variables and their defaults
type Config struct {
	App    *App
	Logger *Logger
	Server *Server
	JWT    *JWT
}

// ParseConfigs parses configs from env variables to fill Config
func ParseConfigs() (*Config, error) {
	c := &Config{
		App:    &App{},
		Logger: &Logger{},
		Server: &Server{},
		JWT:    &JWT{},
	}

	if err := env.Parse(c); err != nil {
		return nil, err
	}

	return c, nil
}

// lookatme
// App meta info about your App
type App struct {
	Name      string `env:"APP_NAME" envDefault:"App"`
	Version   string `env:"APP_VERSION" envDefault:"v0.0.0"`
	GoVersion string `env:"GO_VERSION" envDefault:"0.0.0"`
}

// lookatme
// Logger things to describe how you expect your logger to work
type Logger struct {
	Level      int  `env:"LOG_LEVEL" envDefault:"-1"`
	Structured bool `env:"LOG_STRUCTURED" envDefault:"false"`
}

// lookatme
// Server consider adding things like timeout and make sure this matches with the Dockerfile
type Server struct {
	Host string `env:"SERVER_HOST" envDefault:":"`
	Port string `env:"SERVER_PORT" envDefault:"8000"`
}

// lookatme
// JWT make sure to pass something into this before prod
type JWT struct {
	Key string `env:"JWT_KEY" envDefault:"your-256-bit-secret"` // replaceme: hopefully this isn't your key
}
