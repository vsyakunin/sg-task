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

var Messages = []Message{MessageOne, MessageTwo, MessageThree, MessageFour}

var MessageOne = Message{
	ID:         1,
	TaskID:     1,
	FromUserID: 1,
	Text:       "message one",
	FileKey:    "fileone.txt",
	Created:    time.Now(),
}

var MessageTwo = Message{
	ID:         2,
	TaskID:     1,
	FromUserID: 2,
	Text:       "message two",
	FileKey:    "filetwo.txt",
	Created:    time.Now(),
}

var MessageThree = Message{
	ID:         3,
	TaskID:     2,
	FromUserID: 3,
	Text:       "message three",
	FileKey:    "filethree.txt",
	Created:    time.Now(),
}

var MessageFour = Message{
	ID:         4,
	TaskID:     2,
	FromUserID: 3,
	Text:       "message four",
	FileKey:    "filefour.txt",
	Created:    time.Now(),
}
