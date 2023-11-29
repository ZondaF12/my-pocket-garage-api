package activity

import (
	"net/http"

	"github.com/ZondaF12/my-pocket-garage/internal/database"
	"github.com/gofiber/fiber/v2"
)

// @Summary Add activity to a vehicle
// @Description add activity to a vehicles database.
// @Tags Activity
// @Accept */*
// @Produce plain
// @Success 200 {object} database.Activity
// @Router /api/user/:userId/vehicles/:vehicleReg/activity [post]
func HandleAddVehicleActivity(c *fiber.Ctx) error {
	payload := database.Activity{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": err,
		})
	}

	err := database.AddVehicleActivity(payload)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Sucessfully added activity",
	})
}