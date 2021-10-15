package main

import (
	"matcha/install"
	"matcha/utils/dbcon"
	"matcha/utils/logger"

	"github.com/savsgio/atreugo/v11"
)

func main() {
	// Closing database connection if
	defer dbcon.Close()

	// Application installation
	if err := install.InstallMatchaApplication(dbcon.Get(), logger.Get()); err != nil {
		logger.Get().Fatalf("fatal: InstallMatchaApplication(): %v", err)
	}

	// Server configuration
	app := atreugo.New(atreugo.Config{
		Addr:   ":80",
		Name:   "golang/1.17/atreugo",
		Logger: logger.Get(),
	})

	SetupApiV1Router(app)

	// Launching application's listener
	if err := app.ListenAndServe(); err != nil {
		logger.Get().Fatalf("fatal: ListenAndServe(): %v", err)
	}
}
