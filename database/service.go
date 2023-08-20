package database

import (
	"blink-go-gin-boilerplate/migration"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Initialize(config *Config) (*gorm.DB, error) {
	dsn := config.generateDSN()

	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), 
		logger.Config{
			SlowThreshold:             time.Second,   
			LogLevel:                  logger.Silent, 
			IgnoreRecordNotFoundError: true,          
			ParameterizedQueries:      true,          
			Colorful:                  false,         
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger,
	})

	migration.Migrate(db)

	return db, err
}