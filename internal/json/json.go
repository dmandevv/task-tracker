package json

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/dmandevv/task-tracker/internal/config"
)

func SaveTasksToFile(cfg *config.Config) error {
	jsonData, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("Error marshaling to JSON: %v", err)
	}

	file, err := os.Create(getFilePath())
	if err != nil {
		return fmt.Errorf("Error creating json file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(jsonData)
	if err != nil {
		return fmt.Errorf("Error writing to json file: %v", err)
	}

	return nil
}

func LoadTasksFromFile() (*config.Config, error) {
	file, err := os.Open(getFilePath())
	if err != nil {
		return nil, fmt.Errorf("Error opening json file: %v", err)
	}
	defer file.Close()

	var cfg config.Config
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&cfg)
	if err != nil {
		return nil, fmt.Errorf("Error decoding json file: %v", err)
	}

	return &cfg, nil
}

func getFilePath() string {
	saveFilePath := os.Getenv("TASK_TRACKER_SAVE_FILE_PATH")
	if saveFilePath == "" {
		fmt.Println("Environment variable TASK_TRACKER_SAVE_FILE_PATH is not set or empty, using default 'mytasks.json'.")
		return "./tasks.json"
	} else {
		fmt.Printf("Using save file path from environment: %s\n", saveFilePath)
		return saveFilePath
	}
}
