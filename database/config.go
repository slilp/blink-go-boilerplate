package database

import "fmt"

type Config struct {
	PostgresUser     string `envconfig:"POSTGRES_USER" default:"pquser"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" default:"yourpassword"`
	PostgresHost     string `envconfig:"POSTGRES_HOST" default:"localhost"` 
	PostgresPort     string `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresDB       string `envconfig:"POSTGRES_DB" default:"yourdb"`
	PostgresMode     string `envconfig:"POSTGRES_MODE" default:"disable"`
}

func (c *Config) generateDSN() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=%s",
		c.PostgresHost, c.PostgresUser, c.PostgresPassword, c.PostgresDB, c.PostgresPort,c.PostgresMode)

	return dsn
}

