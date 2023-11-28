package handlers

import "github.com/gofiber/fiber/v2"

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

// @Summary Login user.
// @Description logs in the user.
// @Tags user
// @Accept */*
// @Produce plain
// @Success 200 "You are logged in"
// @Router /user [get]
func (u *UserController) Profile(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "You are logged in",
	})
}

