package models

import "time"

type Task struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Name      string    `json:"name"`
	Category  string    `json:"category"`
	Status    string    `json:"status"`
	StartedAt time.Time `json:"started_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

var AllTasks = []Task{TaskOne, TaskTwo, TaskThree}

var TaskOne = Task{
	ID:        1,
	UserID:    1,
	Name:      "task one",
	Category:  "find a doctor",
	Status:    "processing",
	StartedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var TaskTwo = Task{
	ID:        2,
	UserID:    1,
	Name:      "task two",
	Category:  "buy plane tickets",
	Status:    "processing",
	StartedAt: time.Now(),
	UpdatedAt: time.Now(),
}

var TaskThree = Task{
	ID:        3,
	UserID:    2,
	Name:      "task three",
	Category:  "find a dog sitter",
	Status:    "processing",
	StartedAt: time.Now(),
	UpdatedAt: time.Now(),
}
