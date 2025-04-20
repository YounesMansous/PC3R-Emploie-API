package models

import "time"

type EventsJSON struct {
	ID      int64     `json:"id"`
	Titre   string    `json:"titre"`
	Message string    `json:"message"`
	Date    time.Time `json:"created_at"`
}
