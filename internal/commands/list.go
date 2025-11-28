package commands

import (
	"github.com/dmandevv/task-tracker/internal/config"
	"github.com/dmandevv/task-tracker/internal/task"
)

func GetTasksByFilter(cfg *config.Config, status task.Status) []task.Task {
	filteredTasks := make([]task.Task, 0)
	for _, t := range cfg.Tasks {
		if t.Status == status {
			filteredTasks = append(filteredTasks, t)
		}
	}
	return filteredTasks
}
