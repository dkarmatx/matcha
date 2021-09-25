package debug_api

import (
	"encoding/json"
	"fmt"
	"matcha/models"
	"time"

	"github.com/savsgio/atreugo/v11"
)

func HandleUserCreate(ctx *atreugo.RequestCtx) error {
	resp := ResponseStatus{Success: true, RequestedAt: time.Now()}
	user := models.User{}

	// fill user from body JSON
	if err := json.Unmarshal(ctx.PostBody(), &user); err == nil {
		if err := user.Insert(); err != nil {
			resp.Success = false
			resp.Message = fmt.Sprintf("db_error: %v", err)
		}
	} else {
		resp.Success = false
		resp.Message = fmt.Sprintf("json_decode_error: %v", err)
	}
	return ctx.JSONResponse(resp)
}

// func HandleUserListing(ctx *atreugo.RequestCtx) error {
// 	resp := ResponseStatus{Success: true, Message: "", RequestedAt: time.Now()}
// 	_, err := dbcon.Get().Query(`SELECT * FROM users`)

// 	if err != nil {
// 		resp.Message = "" + err.Error()
// 		resp.Success = false
// 	}
// 	return ctx.JSONResponse(resp)
// }

// func HandleUserRemove(ctx *atreugo.RequestCtx) error {
// 	resp := ResponseStatus{Success: true, Message: "", RequestedAt: time.Now()}
// 	user := models.User{}

// 	_, err := dbcon.Get().Query(`
// 		DELETE FROM users
// 		WHERE user_id = $1`, user.Id)
// 	if err != nil {
// 		resp.Message = "" + err.Error()
// 		resp.Success = false
// 	}
// 	return ctx.JSONResponse(resp)
// }
