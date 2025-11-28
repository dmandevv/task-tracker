package commands

import (
	"fmt"

	"github.com/dmandevv/task-tracker/internal/config"
)

func DeleteTask(cfg *config.Config, id int) error {
	for index, task := range cfg.Tasks {
		if task.ID == id {
			cfg.Tasks = append(cfg.Tasks[:index], cfg.Tasks[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("task with ID \"%d\" not found", id)
}
