package handler

import (
	"strconv"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"

	"github.com/yourusername/go-user-api/internal/models"
	"github.com/yourusername/go-user-api/internal/service"
)

// UserHandler holds the service dependency and validator instance.
type UserHandler struct {
	svc      service.UserService
	validate *validator.Validate
}

// NewUserHandler creates a UserHandler with a pre-configured validator
// that includes the custom "dob_format" validation rule.
func NewUserHandler(svc service.UserService) *UserHandler {
	v := validator.New()

	// Register custom validator: dob must be in YYYY-MM-DD format
	_ = v.RegisterValidation("dob_format", func(fl validator.FieldLevel) bool {
		_, err := time.Parse("2006-01-02", fl.Field().String())
		return err == nil
	})

	return &UserHandler{svc: svc, validate: v}
}

// CreateUser handles POST /users
// Creates a new user and returns 201 with id, name, dob.
func (h *UserHandler) CreateUser(c *fiber.Ctx) error {
	var req models.CreateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "invalid request"})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "invalid request"})
	}

	resp, err := h.svc.CreateUser(c.Context(), req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "internal server error"})
	}

	return c.Status(fiber.StatusCreated).JSON(resp)
}

// GetUser handles GET /users/:id
// Returns the user with dynamically calculated age field.
func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "invalid request"})
	}

	resp, err := h.svc.GetUser(c.Context(), int32(id))
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "user not found"})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

// ListUsers handles GET /users
// Supports optional pagination: ?page=1&limit=10
func (h *UserHandler) ListUsers(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}

	resp, err := h.svc.ListUsers(c.Context(), int32(page), int32(limit))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(models.ErrorResponse{Error: "internal server error"})
	}

	// Return empty array, never null
	if resp == nil {
		resp = []models.UserResponse{}
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

// UpdateUser handles PUT /users/:id
// Updates name and dob, returns updated user without age field.
func (h *UserHandler) UpdateUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "invalid request"})
	}

	var req models.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "invalid request"})
	}

	if err := h.validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "invalid request"})
	}

	resp, err := h.svc.UpdateUser(c.Context(), int32(id), req)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "user not found"})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}

// DeleteUser handles DELETE /users/:id
// Deletes the user and returns 204 No Content.
func (h *UserHandler) DeleteUser(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil || id <= 0 {
		return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "invalid request"})
	}

	if err := h.svc.DeleteUser(c.Context(), int32(id)); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "user not found"})
	}

	return c.SendStatus(fiber.StatusNoContent)
}
