package model

import "time"

type TaskPriority struct {
	High   string
	Medium string
	Low    string
}

var Priority = TaskPriority{
	High:   "High",
	Medium: "Medium",
	Low:    "Low",
}

type Task struct {
	Id        int          `db:"id"`
	Task      string       `db:"task"`
	Priority  TaskPriority `db:"priority"`
	Status    bool         `db:"status"`
	CreatedAt time.Time    `db:"created_at"`
	UserId    int          `db:"user_id"`
}
