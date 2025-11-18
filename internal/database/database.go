package database

import (
	"context"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

// Pool is the global connection pool used across the application.
var Pool *pgxpool.Pool

// Connect establishes a connection to Postgres using pgxpool and
// applies the initial migrations for the Notes API template.
func Connect() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal().Msg("❌ DATABASE_URL is not set")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Unable to create database connection pool")
	}

	if err := pool.Ping(ctx); err != nil {
		log.Fatal().Err(err).Msg("❌ Database ping failed")
	}

	Pool = pool

	// Run migration for notes table
	_, err = pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS notes (
			id SERIAL PRIMARY KEY,
			title TEXT NOT NULL,
			content TEXT,
			created_at TIMESTAMP DEFAULT NOW()
		);
	`)
	if err != nil {
		log.Fatal().Err(err).Msg("❌ Migration failed")
	}

	log.Info().Msg("✅ Migration applied (notes table)")
	log.Info().Msg("✅ Database connected successfully")
}
