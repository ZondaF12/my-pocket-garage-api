package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// @Description get the status of server.
// @Tags health
// @Accept */*
// @Produce plain
// @Success 200 "Hello World"
// @Router /health [get]
func HandleHealthCheck(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Hello World",
	})
}