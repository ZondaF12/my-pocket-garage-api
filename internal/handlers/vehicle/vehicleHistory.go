package vehicle

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ZondaF12/my-pocket-garage/internal/database"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get all of the vehicles to do with a user.
// @Description get the vehicles of a user.
// @Tags User Vehicles
// @Accept */*
// @Produce plain
// @Success 200 {object} database.UserVehicle
// @Router /api/user/:userId/vehicleHistory [get]
func HandlerGetUserVehicleHistory(c *fiber.Ctx) error {
	userId, err := url.QueryUnescape(c.Params("userId"))
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprint(err),
		})
	}

	res, err := database.GetUserVehicleHistory(userId)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprint(err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}