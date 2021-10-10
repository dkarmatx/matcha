package api

import "time"

type ResponseStatus struct {
	Success    bool        `json:"success"`
	Message    string      `json:"message,omitempty"` // it could be omitted if success field is true
	FinishedAt time.Time   `json:"finished_at"`       // time when we started to proccessing the request
	Obj        interface{} `json:"obj,omitempty"`
}

func NewResponseStatus(success bool, msg string, obj interface{}) ResponseStatus {
	return ResponseStatus{Success: success, Message: msg, FinishedAt: time.Now(), Obj: obj}
}
