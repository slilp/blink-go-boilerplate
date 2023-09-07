package app

import (
	"time"

	"github.com/slilp/blink-go-boilerplate/database"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Mode    string   `envconfig:"HTTP_MODE" default:"debug"`
	LogSkip []string `envconfig:"HTTP_LOG_SKIP" default:"/health"`
	Port    int      `envconfig:"HTTP_PORT" default:"5000"`
	Prefix  string   `envconfig:"HTTP_PATH_PREFIX" default:""`

	AllowedOrigins []string      `envconfig:"HTTP_ALLOWED_ORIGINS" default:"*"`
	AllowedHeaders []string      `envconfig:"HTTP_ALLOWED_HEADERS" default:"*"`
	CORSMaxAge     time.Duration `envconfig:"HTTP_CORS_MAX_AGE" default:"24h"`
	Database *database.Config
}


func loadConfig() (*Config, error) {
	var config Config

	if err := envconfig.Process("", &config); err != nil {
		return nil, err
	}

	return &config, nil
}
