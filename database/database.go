package database

import (
	"context"
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

func match1() *model.Match {
	minParticipants := int32(2)
	maxParticipants := int32(4)
	createData := model.MatchCreate{
		Sport:             "Tennis",
		MinParticipants:   &minParticipants,
		MaxParticipants:   &maxParticipants,
		Location:          "Local Tennis Court",
		Description:       "Welcome to this awesome tennis match!",
		ParticipationFee:  0,
		RequiredEquipment: []string{"Racket", "Shoes"},
		Level:             "Any",
		ChatLink:          "https://example.com",
		StartsAt:          time.Date(2025, time.January, 1, 10, 0, 0, 0, time.UTC),
		EndsAt:            time.Date(2025, time.January, 1, 11, 30, 0, 0, time.UTC),
	}

	dbMatch := createData.Match()
	dbMatch.HostUserID = "DemoUser"

	return &dbMatch
}

func match2() *model.Match {
	minParticipants := int32(2)
	maxParticipants := int32(4)
	createData := model.MatchCreate{
		Sport:             "Badminton",
		MinParticipants:   &minParticipants,
		MaxParticipants:   &maxParticipants,
		Location:          "Sports Hall, Downtown",
		Description:       "Looking for people to play badminton with :)",
		ParticipationFee:  1000,
		RequiredEquipment: []string{},
		Level:             "Any",
		ChatLink:          "https://example.com",
		StartsAt:          time.Date(2025, time.January, 10, 18, 0, 0, 0, time.UTC),
		EndsAt:            time.Date(2025, time.January, 10, 19, 0, 0, 0, time.UTC),
	}

	dbMatch := createData.Match()
	dbMatch.HostUserID = "DemoUser"

	return &dbMatch
}

func createInitialData() {
	ctx := context.TODO()

	m := dal.Q.Match
	p := dal.Q.Participation

	dbMatch1 := match1()
	dbMatch2 := match2()

	if err := m.WithContext(ctx).Create(dbMatch1, dbMatch2); err != nil {
		stdlog.Fatalf("Failed to create initial data: %s", err.Error())
	}

	dbParticipation1 := model.Participation{MatchID: dbMatch1.ID, UserID: dbMatch1.HostUserID}
	if err := p.WithContext(ctx).Create(&dbParticipation1); err != nil {
		stdlog.Fatalf("Failed to create initial data: %s", err.Error())
	}

	dbParticipation2 := model.Participation{MatchID: dbMatch2.ID, UserID: dbMatch2.HostUserID}
	if err := p.WithContext(ctx).Create(&dbParticipation2); err != nil {
		stdlog.Fatalf("Failed to create initial data: %s", err.Error())
	}

	stdlog.Println("Data created")
}

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
