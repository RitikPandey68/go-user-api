package middleware

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// RequestDuration is a Fiber middleware that measures and logs the duration
// of every HTTP request using Uber Zap in the format:
//
//	GET /users/1 completed in 4ms
func RequestDuration(logger *zap.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()

		// Process the request chain
		err := c.Next()

		duration := time.Since(start)

		logger.Info("request completed",
			zap.String("method", c.Method()),
			zap.String("path", c.Path()),
			zap.Int("status", c.Response().StatusCode()),
			zap.Duration("duration", duration),
			zap.String("request_id", c.Get("X-Request-ID")),
		)

		return err
	}
}
