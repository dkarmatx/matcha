package debug_api

import (
	"encoding/json"
	"fmt"
	"matcha/models"
	"strconv"
	"time"

	"github.com/savsgio/atreugo/v11"
)

func HandleUserCreate(ctx *atreugo.RequestCtx) error {
	resp := ResponseStatus{Success: true, RequestedAt: time.Now()}
	user := models.User{}

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

type responseUserListing struct {
	ResponseStatus
	Users models.UserList `json:"users,omitempty"`
}

func HandleUserListing(ctx *atreugo.RequestCtx) error {
	var users models.UserList
	resp := responseUserListing{
		ResponseStatus: ResponseStatus{Success: true, RequestedAt: time.Now()},
		Users:          nil,
	}
	err := users.SelectAll()

	if err != nil {
		resp.Message = fmt.Sprintf("db_error: %v", err)
		resp.Success = false
		resp.Users = nil
	} else {
		resp.Users = users
	}
	return ctx.JSONResponse(resp)
}

func HandleUserDelete(ctx *atreugo.RequestCtx) error {
	resp := ResponseStatus{Success: true, RequestedAt: time.Now()}
	id_str := ctx.UserValue("id").(string)

	if id, err := strconv.ParseInt(id_str, 10, 64); err == nil {
		user := models.User{Id: id}
		if err = user.DeleteById(); err != nil {
			resp.Success = false
			resp.Message = fmt.Sprintf("db_error: %v", err)
		}
	} else {
		resp.Success = false
		resp.Message = fmt.Sprintf("user_param: %v", err)
	}

	return ctx.JSONResponse(resp)
}
