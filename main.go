package main

import (
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"notes-memory-core/internal/database"
	"notes-memory-core/internal/handlers"
	"notes-memory-core/internal/middleware"
)

func main() {
	// Pretty logger output during development
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})

	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Warn().Msg("No .env file found, using system environment variables")
	}

	// Connect to Postgres
	database.Connect()

	// Create Fiber app
	app := fiber.New()

	// Middleware
	app.Use(middleware.MetricsMiddleware)
	app.Use(middleware.LoggerMiddleware)

	// Routes
	app.Get("/health", handlers.HealthCheck)
	app.Get("/notes", handlers.GetNotes)
	app.Post("/notes", handlers.CreateNote)
	app.Get("/metrics", func(c *fiber.Ctx) error {
		return c.JSON(middleware.GetMetrics())
	})

	// Port configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Info().
		Str("port", port).
		Msg("🚀 Starting server")

	if err := app.Listen(":" + port); err != nil {
		log.Fatal().Err(err).Msg("❌ Failed to start server")
	}
}
