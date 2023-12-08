package vehicle

import (
	"net/http"

	"github.com/ZondaF12/my-pocket-garage/internal/handlers/tools"
	"github.com/gofiber/fiber/v2"
)

// @Summary Get DVLA info on a specified vehicle.
// @Description get the DVLA data on a vehicle.
// @Tags Vehicle Info
// @Accept */*
// @Produce plain
// @Success 200 {object} tools.VehicleData
// @Router /api/vehicle/:vehicleReg/info [get]
func HandleVehicleInfo(c *fiber.Ctx) error {
	res, err := tools.DoVehicleInfoRequest(c.Params("vehicleReg"))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": err,
		})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": res,
	})
}
