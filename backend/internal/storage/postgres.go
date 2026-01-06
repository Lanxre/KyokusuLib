package storage

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/lanxre/kyokusulib/internal/config" 
	"github.com/pressly/goose/v3"
	"go.uber.org/fx"
)

const (
	DefaultTimeout = 5 * time.Second
)

func NewPostgresDB(lc fx.Lifecycle, cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, fmt.Errorf("failed to open db connection: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), DefaultTimeout)
	defer cancel()

	if err := db.PingContext(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	if err := runMigrations(db); err != nil {
		return nil, fmt.Errorf("migration failed: %w", err)
	}

	if err := runSeeds(db); err != nil {
		return nil, fmt.Errorf("seeding failed: %w", err)
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Println("Closing database connection...")
			return db.Close()
		},
	})

	return db, nil
}

func runMigrations(db *sql.DB) error {
	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return fmt.Errorf("could not run migrations: %w", err)
	}

	log.Println("Database migrations check completed")
	return nil
}

func runSeeds(db *sql.DB) error {
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

	levelDefinitionsQuery := `
		INSERT INTO level_definitions (level, title, total_xp_required)
		VALUES
			(1, 'Новичок', 0),
			(2, 'Ученик', 1000)
		ON CONFLICT (level) DO NOTHING;
	`

	if _, err := db.ExecContext(ctx, levelDefinitionsQuery); err != nil {
		log.Printf("Warning: failed to seed level definitions: %v", err)
	}
	
	return nil
}