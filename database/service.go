package database

import (
	"log"
	"os"
	"time"

	"github.com/slilp/blink-go-boilerplate/migration"

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