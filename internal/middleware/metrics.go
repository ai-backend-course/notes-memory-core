package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// Simple in-memory metrics for API monitoring.
// This tracks total requests, errors, and average latency.
var (
	TotalRequests  int
	TotalErrors    int
	TotalLatencyMs float64
)

// MetricsMiddleware measures request duration and updates counters.
func MetricsMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	err := c.Next()

	elapsed := time.Since(start).Milliseconds()

	TotalRequests++
	TotalLatencyMs += float64(elapsed)

	if err != nil {
		TotalErrors++
	}

	return err
}

// GetMetrics returns aggregated metrics as JSON.
func GetMetrics() fiber.Map {
	return fiber.Map{
		"total_requests": TotalRequests,
		"total_errors":   TotalErrors,
		"avg_latency_ms": averageLatency(),
	}
}

// averageLatency computes the average request latency in ms.
func averageLatency() float64 {
	if TotalRequests == 0 {
		return 0
	}
	return TotalLatencyMs / float64(TotalRequests)
}
