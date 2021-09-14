package main

import (
	"database/sql"
	"log"
	"matcha/config"

	_ "github.com/lib/pq"
	"github.com/savsgio/atreugo/v11"
)

func main() {
	app := atreugo.New(atreugo.Config{
		Addr: ":80",
		Name: "golang/1.17/atreugo",
	})

	log.Println("DSN: ", config.GetDSN())

	app.GET("/ping_db", pingDB)

	if err := app.ListenAndServe(); err != nil {
		panic(err)
	}
}

func pingDB(ctx *atreugo.RequestCtx) error {
	var db *sql.DB
	var err error
	var message string

	if db, err = sql.Open("postgres", config.GetDSN()); err == nil {
		defer db.Close()

		if err = db.Ping(); err == nil {
			message = "OK"
		}
	}

	return ctx.JSONResponse(atreugo.JSON{
		"message": message,
		"error":   err,
	})
}
