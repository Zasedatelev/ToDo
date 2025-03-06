package model

import "time"

type Task struct {
	TaskId      int32     `json: "task_id"`
	Title       string    `json: "title"`
	Description string    `json: "description"`
	Status      string    `json: "status"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
