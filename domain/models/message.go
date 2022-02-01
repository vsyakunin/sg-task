package models

import "time"

type Message struct {
	ID         int64     `json:"id"`
	TaskID     int64     `json:"task_id"`
	FromUserID int64     `json:"from_user_id"`
	Text       string    `json:"message_text"`
	FileKey    string    `json:"file_key"`
	Created    time.Time `json:"created_at"`
}
