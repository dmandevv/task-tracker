package task

import "time"

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type Status int

const (
	TODO Status = iota
	IN_PROGRESS
	DONE
)

func (s Status) String() string {
	switch s {
	case TODO:
		return "todo"
	case IN_PROGRESS:
		return "in-progress"
	case DONE:
		return "done"
	default:
		return "UNKNOWN"
	}
}
