package debug_api

import (
	"matcha/dbcon"

	"github.com/savsgio/atreugo/v11"
)

// GET /debug/db/ping - ping database service
// Return JSON with params:
// - success: bool		- true if db service is avaible and Ping is proprely finished
// - message: string	- message related to status
func handlePingDB(ctx *atreugo.RequestCtx) error {
	message := "db_service: avaible"
	err := dbcon.Get().Ping()
	if err != nil {
		message = "db_service: unavaible due error: " + err.Error()
	}

	return ctx.JSONResponse(atreugo.JSON{
		"message": message,
		"success": err == nil,
	})
}

func SetHandlers(app *atreugo.Atreugo) {
	debug_router := app.NewGroupPath("/debug")

	debug_router.GET("/db/ping", handlePingDB)
}
