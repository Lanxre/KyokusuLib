package storage

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func NewPostgresDB(databaseURL string) *sql.DB {
	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatal("Failed to open db connection:", err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping db:", err)
	}

	runMigrations(db)

	return db
}

func runMigrations(db *sql.DB) {
	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatal("Failed to set goose dialect:", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		log.Fatal("Could not run migrations:", err)
	}

	log.Println("Database migrations check completed")
}
