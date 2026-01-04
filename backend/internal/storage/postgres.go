package storage

import (
	"context"
	"database/sql"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

const (
	DefaultTimeout = 5 * time.Second
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
	
	runSeeds(db)
	
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

func runSeeds(db *sql.DB) {
	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	log.Println("Checking seed data...")
	
	tagsQuery := `
		INSERT INTO user_tags (tag)
		VALUES 
			('Читатель')
		ON CONFLICT (tag) DO NOTHING;
	`
	if _, err := db.ExecContext(ctx, tagsQuery); err != nil {
		log.Printf("Warning: failed to seed tags: %v", err)
	}
	
	levelDefinitionsQuery :=`
		INSERT INTO level_definitions (level, title, total_xp_required)
		VALUES
			(1, 'Новичок', 0),
			(2, 'Ученик', 1000)
		ON CONFLICT (level) DO NOTHING;
	`
	
	if _, err := db.ExecContext(ctx, levelDefinitionsQuery); err != nil {
		log.Printf("Warning: failed to seed level definitions: %v", err)
	}
}

