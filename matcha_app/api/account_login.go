package api

import (
	"encoding/json"
	"fmt"
	"matcha/auth"
	"matcha/models"
	"time"

	"github.com/savsgio/atreugo/v11"
)

type accountLoginRequest struct {
	Name string `json:"user_nickname"`
	Pass string `json:"user_password"`
}

func HandleAccountLogin(ctx *atreugo.RequestCtx) error {
	success := true
	message := ""

	body := ctx.Request.Body()

	var account accountLoginRequest
	if err := json.Unmarshal(body, &account); err != nil {
		message = "invalid json body"
		return ctx.JSONResponse(NewResponseStatus(false, message, nil), 400)
	}

	var ulst models.UserList
	if err := ulst.SelectByName(account.Name); err != nil {
		message = fmt.Sprintf("database error: %v", err)
		return ctx.JSONResponse(NewResponseStatus(false, message, nil), 500)
	}

	if len(ulst) == 0 {
		message = fmt.Sprintf("user not found: \"%s\"", account.Name)
		return ctx.JSONResponse(NewResponseStatus(false, message, nil), 404)
	}

	u := ulst[0]
	if string(u.PassHash) != string(auth.CalcHashPassword([]byte(account.Pass), u.PassSalt)) {
		message = "wrong password"
		return ctx.JSONResponse(NewResponseStatus(false, message, nil), 403)
	}

	var token models.AuthToken
	var toklst models.AuthTokenList
	if err := toklst.SelectByUserId(u.Id); err != nil {
		message = fmt.Sprintf("database error: %v", err)
		return ctx.JSONResponse(NewResponseStatus(false, message, nil), 500)
	}

	weekdur := time.Hour * 24 * 7
	if len(toklst) == 0 {
		token = models.AuthTokenCreate(u.Id, weekdur)
	} else {
		token = toklst[0]
		token.ExpiresAt = time.Now().Add(weekdur)
		token.UpdateByValue()
	}

	return ctx.JSONResponse(NewResponseStatus(success, message, token))
}
