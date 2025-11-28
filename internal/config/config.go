package config

import (
	"github.com/dmandevv/task-tracker/internal/task"
)

type Config struct {
	Tasks  []task.Task `json:"tasks"`
	NextID int         `json:"next_id"`
}
