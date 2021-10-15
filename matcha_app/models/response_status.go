package models

import "time"

type ResponseStatus struct {
	Message    string    `json:"message,omitempty"`
	Success    bool      `json:"success"`
	FinishedAt time.Time `json:"finished_at"`
}

func MakeResponseStatus(success bool, message string) ResponseStatus {
	return ResponseStatus{
		Message:    message,
		Success:    success,
		FinishedAt: time.Now(),
	}
}

func (r ResponseStatus) Finish() ResponseStatus {
	r.FinishedAt = time.Now()
	return r
}
