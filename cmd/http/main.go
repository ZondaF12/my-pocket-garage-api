package main

import (
	"github.com/ZondaF12/my-pocket-garage/pkg/app"
	"github.com/ZondaF12/my-pocket-garage/pkg/shutdown"
)

// @title My Pocket Garage API
// @version 1.0
// @description Pocket Garage API with Golang using Fiber and MongoDB
// @contact.name Ruaridh Bell
// @license.name MIT
// @BasePath /
func main() {
	// setup and run app
	cleanup, err := app.SetupAndRunApp()
	if err != nil {
		panic(err)
	}

	defer cleanup()
	
	// ensure the server is shutdown gracefully & app runs
	shutdown.Gracefully()
}
