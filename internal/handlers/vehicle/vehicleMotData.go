package vehicle

import (
	"net/http"

	"github.com/ZondaF12/my-pocket-garage/internal/handlers/tools"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get DVSA mot data on a specified vehicle.
// @Description get the DVSA mot data on a vehicle.
// @Tags Vehicle Info
// @Accept */*
// @Produce plain
// @Success 200 {object} tools.MotData
// @Router /api/vehicle/mot/:vehicleReg [get]
func HandleVehicleMotData(c *fiber.Ctx) error {
	res, err := tools.DoVehicleMotRequest(c.Params("vehicleReg"))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}