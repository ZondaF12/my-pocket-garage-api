package vehicle

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/ZondaF12/my-pocket-garage/internal/database"
	"github.com/gofiber/fiber/v2"
)

type AddVehicleBody struct {
	Registration string
}

// @Summary Adds specified vehicle to the userVehicles collection.
// @Description adds specified vehicle to the collection.
// @Tags User Vehicles
// @Accept */*
// @Produce plain
// @Success 200 "Vehicle Added Successfully"
// @Router /api/user/:userId/vehicles [post]
func HandleAddUserVehicle(c *fiber.Ctx) error {
	payload := AddVehicleBody{}

	if err := c.BodyParser(&payload); err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": err,
		})
	}

	// TODO: Validate passed in userId against the one in the
	userId, err := url.QueryUnescape(c.Params("userId"))
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprint(err),
		})
	}

	err = database.AddUserVehicle(userId, payload.Registration)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprint(err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "Vehicle Added Successfully",
	})
}

// @Summary Get all of the vehicles to do with a user.
// @Description get the vehicles of a user.
// @Tags User Vehicles
// @Accept */*
// @Produce plain
// @Success 200 {object} database.UserVehicle
// @Router /api/user/:userId/vehicles [get]
func HandleGetUserVehicles(c *fiber.Ctx) error {
	userId, err := url.QueryUnescape(c.Params("userId"))
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprint(err),
		})
	}

	res, err := database.GetUserVehicles(userId)
	if err != nil {
		return c.Status(http.StatusForbidden).JSON(fiber.Map{
			"error": fmt.Sprint(err),
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}
