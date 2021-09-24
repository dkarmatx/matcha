package debug_api

import (
	"matcha/dbcon"
	"time"

	"github.com/savsgio/atreugo/v11"
)

type ResponseStatus struct {
	Success     bool      `json:"success"`
	Message     string    `json:"message,omitempty"` // it could be omitted if success field is true
	RequestedAt time.Time `json:"requested_at"`      // time when we started to proccessing the request
}

// GET /debug/db/ping - ping database service
// Return JSON with params:
// - success: bool		- true if db service is avaible and Ping is proprely finished
// - message: string	- message related to status
func handlePingDB(ctx *atreugo.RequestCtx) error {
	resp := ResponseStatus{Success: true, Message: "", RequestedAt: time.Now()}

	err := dbcon.Get().Ping()
	if err != nil {
		resp.Message = "db_service: unavaible due error: " + err.Error()
		resp.Success = false
	}

	return ctx.JSONResponse(resp)
}

func SetHandlers(app *atreugo.Atreugo) {
	debug_router := app.NewGroupPath("/debug")

	debug_router.GET("/ping_database", handlePingDB)
}
