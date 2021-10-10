package main

import (
	"matcha/api"
	"matcha/dbcon"
	"matcha/debug_api"
	"matcha/install"
	"matcha/logger"

	"github.com/savsgio/atreugo/v11"
)

func main() {
	// Closing database connection if
	defer dbcon.Close()

	// Application installation
	if err := install.InstallMatchaApplication(dbcon.Get(), logger.Get()); err != nil {
		logger.Get().Fatalf("Exiting due fatal error has occured: Install(): %s", err.Error())
	}

	// Server configuration
	app := atreugo.New(atreugo.Config{
		Addr:   ":80",
		Name:   "golang/1.17/atreugo",
		Logger: logger.Get(),
	})

	// Routing: `/debug/*`
	debug_api.SetHandlers(app)
	// Routing: `/api/v1/*`
	api.SetupRouters(app)

	// Launching application's listener
	if err := app.ListenAndServe(); err != nil {
		logger.Get().Fatalf("Exiting due fatal error has occured: ListenAndServe(): %s", err.Error())
	}
}
