package debug_api

import (
	"matcha/dbcon"
	"time"

	"github.com/savsgio/atreugo/v11"
)

func HandlePingDB(ctx *atreugo.RequestCtx) error {
	resp := ResponseStatus{Success: true, Message: "", RequestedAt: time.Now()}

	err := dbcon.Get().Ping()
	if err != nil {
		resp.Message = "db_service: unavaible due error: " + err.Error()
		resp.Success = false
	}

	return ctx.JSONResponse(resp)
}
