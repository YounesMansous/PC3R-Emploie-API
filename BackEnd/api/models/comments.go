package models

import "time"

type Comments struct {
	ID      int8      `json:"id"`
	Content string    `json:"content"`
	Date    time.Time `json:"created_at"`
	Event   int8      `json:"event_id"`
	User    int8      `json:"user_id"`
}

type CommentsJSON struct {
	ID      int8      `json:"id"`
	Content string    `json:"content"`
	User    string    `json:"user"`
	Date    time.Time `json:"created_at"`
}
