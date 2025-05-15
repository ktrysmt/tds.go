package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"tds.go/config"
	"tds.go/pkg/infrastructure/logger"
)

func NewPostgresDB(cfg *config.Config) *sql.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Error("Failed to connect to database",
			logger.Field("error", err.Error()),
		)
		panic(err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		logger.Error("Failed to ping database",
			logger.Field("error", err.Error()),
		)
		panic(err)
	}

	logger.Info("Database connected successfully",
		logger.Field("host", cfg.Database.Host),
		logger.Field("database", cfg.Database.DBName),
	)

	return db
}
