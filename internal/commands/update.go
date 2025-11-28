package commands

import (
	"fmt"
	"time"

	"github.com/dmandevv/task-tracker/internal/config"
	"github.com/dmandevv/task-tracker/internal/task"
)

func UpdateTask(cfg *config.Config, id int, description string) error {
	for index, task := range cfg.Tasks {
		if task.ID == id {
			cfg.Tasks[index].Description = description
			cfg.Tasks[index].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("task with ID \"%d\" not found", id)
}

func MarkTask(cfg *config.Config, id int, status task.Status) error {
	for index, task := range cfg.Tasks {
		if task.ID == id {
			cfg.Tasks[index].Status = status
			cfg.Tasks[index].UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("task with ID \"%d\" not found", id)
}
