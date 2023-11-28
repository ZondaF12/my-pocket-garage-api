package app

import (
	"fmt"

	"github.com/ZondaF12/my-pocket-garage/config"
	"github.com/ZondaF12/my-pocket-garage/internal/database"
	"github.com/ZondaF12/my-pocket-garage/internal/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func SetupAndRunApp() (func(), error) {
	// load env
	env, err := config.LoadConfig()
	if err != nil {
		fmt.Printf("error: %v", err)
		return nil, err
	}

	// start database
	err = database.StartMongoDB(env.MONGODB_URI, env.DATABASE)
	if err != nil {
		fmt.Println("foo")
		return nil, err
	}

	// create app
	app := fiber.New()

	// attach middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// setup routes
	router.SetupRoutes(app, env)

	// get the port and start
	app.Listen(":" + env.PORT)

	// return a function to close the server and database
	return func() {
		database.CloseMongoDB()
		app.Shutdown()
	}, nil
}