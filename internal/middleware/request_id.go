package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// RequestID is a Fiber middleware that injects a unique X-Request-ID header
// into every request and response. If the client already sends one, it is
// preserved; otherwise a fresh UUID v4 is generated.
func RequestID() fiber.Handler {
	return func(c *fiber.Ctx) error {
		requestID := c.Get("X-Request-ID")
		if requestID == "" {
			requestID = uuid.New().String()
		}

		// Expose the ID on both the request context and the response headers
		c.Set("X-Request-ID", requestID)
		c.Locals("requestID", requestID)

		return c.Next()
	}
}
