package database

import "fmt"

type Config struct {
	PostgresUser     string `envconfig:"POSTGRES_USER" default:"blink"`
	PostgresPassword string `envconfig:"POSTGRES_PASSWORD" default:"Ngineer950712"`
	PostgresHost     string `envconfig:"POSTGRES_HOST" default:"localhost"` 
	PostgresPort     string `envconfig:"POSTGRES_PORT" default:"5432"`
	PostgresDB       string `envconfig:"POSTGRES_DB" default:"blink_go"`
}

func (c *Config) generateDSN() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s database=%s port=%s sslmode=require",
		c.PostgresHost, c.PostgresUser, c.PostgresPassword, c.PostgresDB, c.PostgresPort)

	return dsn
}

