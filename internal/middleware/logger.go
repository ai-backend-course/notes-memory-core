package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

// LoggerMiddleware logs each incoming request, including method, route,
// status code, and execution time. This is a lightweight middleware
// suitable for production use with zerolog.
func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next() // process request

	duration := time.Since(start)

	log.Info().
		Str("method", c.Method()).
		Str("path", c.Path()).
		Int("status", c.Response().StatusCode()).
		Dur("duration", duration).
		Msg("Handled request")

	return err
}
