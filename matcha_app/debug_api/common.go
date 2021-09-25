package debug_api

import (
	"time"

	"github.com/savsgio/atreugo/v11"
)

type ResponseStatus struct {
	Success     bool      `json:"success"`
	Message     string    `json:"message,omitempty"` // it could be omitted if success field is true
	RequestedAt time.Time `json:"requested_at"`      // time when we started to proccessing the request
}

func SetHandlers(app *atreugo.Atreugo) {
	debug_router := app.NewGroupPath("/debug")

	debug_router.GET("/ping_database", HandlePingDB)
	debug_router.POST("/user/add", HandleUserCreate)
	// debug_router.GET("/user/list", HandleUserListing)
	// debug_router.DELETE("/user/:id", HandleUserRemove)
}
