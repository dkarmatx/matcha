package main

import (
	"matcha/api/v1"

	"github.com/savsgio/atreugo/v11"
)

func SetupApiV1Router(app *atreugo.Atreugo) {
	rtr := app.NewGroupPath("/api/v1")

	rtr.GET("/health", api.WrapJsonResponse(api.HandleHealth))
}
