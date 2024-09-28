package tasks

import (
	"time"
)

type Task struct {
	Description string
	Status      Status
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type TasksList struct {
	Tasks []Task
}

type Status int8

const (
	InProgress Status = iota
	ToDo       Status = iota
	Done       Status = iota
	All        Status = iota
)

func StatusFromString(status string) Status {
	switch status {
	case "in-progress":
		return InProgress
	case "done":
		return Done
	case "todo":
		return ToDo
	default:
		return All
	}
}

func (s Status) String() string {
	switch s {
	case InProgress:
		return "In Progress"
	case Done:
		return "Done"
	case ToDo:
		return "To Do"
	default:
		return "All"
	}
}
