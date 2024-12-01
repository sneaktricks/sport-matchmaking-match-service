package database

import (
	"fmt"
	"os"
	"time"

	stdlog "log"

	"github.com/sneaktricks/sport-matchmaking-match-service/dal"
	"github.com/sneaktricks/sport-matchmaking-match-service/log"
	"github.com/sneaktricks/sport-matchmaking-match-service/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(&model.Match{}, &model.Participation{})
	log.Logger.Info("Automigration complete")
}

var dbLogger = logger.New(
	stdlog.New(os.Stdout, "\r\n", stdlog.LstdFlags),
	logger.Config{
		LogLevel:             logger.Warn,
		SlowThreshold:        500 * time.Millisecond,
		ParameterizedQueries: true,
	},
)

// Creates a connection to the PostgreSQL database specified by environment variables
// PGUSER, PGPASSWORD, PGHOST, PGPORT, and PGDATABASE.
func Initialize() (db *gorm.DB, err error) {
	user := os.Getenv("PGUSER")
	password := os.Getenv("PGPASSWORD")
	host := os.Getenv("PGHOST")
	port := os.Getenv("PGPORT")
	database := os.Getenv("PGDATABASE")

	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		user,
		password,
		database,
		host,
		port,
	)

	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{TranslateError: true, Logger: dbLogger})
	if err != nil {
		return nil, err
	}
	log.Logger.Info("Successfully connected to database")

	// Configure connection pool limits
	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}
	sqlDB.SetMaxOpenConns(50)
	sqlDB.SetMaxIdleConns(5)

	autoMigrate(db)
	dal.SetDefault(db)
	dal.Use(db)

	// createInitialData()

	return db, nil
}
