package routes

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	"github.com/yourusername/go-user-api/internal/handler"
	"github.com/yourusername/go-user-api/internal/middleware"
)

// RegisterRoutes attaches all application middlewares and routes to the
// provided Fiber app instance. Route layout:
//
//	POST   /users        → CreateUser
//	GET    /users        → ListUsers  (supports ?page=1&limit=10)
//	GET    /users/:id    → GetUser
//	PUT    /users/:id    → UpdateUser
//	DELETE /users/:id    → DeleteUser
func RegisterRoutes(app *fiber.App, h *handler.UserHandler, logger *zap.Logger) {
	// Global middleware
	app.Use(middleware.RequestID())
	app.Use(middleware.RequestDuration(logger))

	// User routes
	users := app.Group("/users")
	users.Post("/", h.CreateUser)
	users.Get("/", h.ListUsers)
	users.Get("/:id", h.GetUser)
	users.Put("/:id", h.UpdateUser)
	users.Delete("/:id", h.DeleteUser)
}
