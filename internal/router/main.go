package router

import (
	"github.com/ZondaF12/my-pocket-garage/config"
	_ "github.com/ZondaF12/my-pocket-garage/docs"
	"github.com/ZondaF12/my-pocket-garage/internal/auth"
	"github.com/ZondaF12/my-pocket-garage/internal/handlers"
	"github.com/ZondaF12/my-pocket-garage/internal/handlers/activity"
	"github.com/ZondaF12/my-pocket-garage/internal/handlers/vehicle"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App, config config.EnvVars) {
	// add docs
	app.Get("/swagger/*", swagger.HandlerDefault)
	
	app.Get("/health", handlers.HandleHealthCheck)

	userGroup := app.Group("/user")

	// middleware to protect routes
	authMiddleware := auth.NewAuthMiddleware(config)
	userGroup.Use(authMiddleware.ValidateToken)
	// auth routes
	userGroup.Get("/", handlers.NewUserController().Profile)

	apiGroup := app.Group("/api")
	// middleware to protect routes
	apiGroup.Use(authMiddleware.ValidateToken)
	// auth routes
	apiGroup.Get("/vehicle/:vehicleReg/info", vehicle.HandleVehicleInfo)
	apiGroup.Get("/vehicle/:vehicleReg/mot", vehicle.HandleVehicleMotData)
	apiGroup.Post("/user/:userId/vehicles/:vehicleReg/activity", activity.HandleAddVehicleActivity)
	apiGroup.Post("/user/:userId/vehicles", vehicle.HandleAddUserVehicle)
	apiGroup.Get("/user/:userId/vehicles", vehicle.HandleGetUserVehicles)
	apiGroup.Get("/user/:userId/activevehicle", vehicle.HandleGetUserVehicles)
}