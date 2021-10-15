package api

import (
	"matcha/models"
	"matcha/utils/dbcon"

	"github.com/savsgio/atreugo/v11"
)

func HandleHealth(ctx *atreugo.RequestCtx) (HttpStatus, AnyJsonObj) {
	rstatus := models.ResponseStatus{Success: true, Message: "Okay"}
	rcode := HttpStatus(200)

	if err := dbcon.Get().Ping(); err != nil {
		rstatus.Message = "db_error: " + err.Error()
		rstatus.Success = false
		rcode = HttpStatus(500)
	}

	return rcode, rstatus.Finish()
}
