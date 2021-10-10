package api

import "github.com/savsgio/atreugo/v11"

func SetupRouters(app *atreugo.Atreugo) {
	api_main_router := app.NewGroupPath("/api/v1")

	api_main_router.POST("/account/login", HandleAccountLogin)
}
