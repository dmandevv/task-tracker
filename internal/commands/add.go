package commands

import (
	"fmt"
	"time"

	"github.com/dmandevv/task-tracker/internal/config"
	"github.com/dmandevv/task-tracker/internal/task"
)

func AddTask(cfg *config.Config, description string) {
	newTask := task.Task{
		ID:          cfg.NextID,
		Description: description,
		Status:      task.TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	cfg.Tasks = append(cfg.Tasks, newTask)
	cfg.NextID++
	fmt.Printf("%v added\n", description)
}
