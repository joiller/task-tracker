package task

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func tasksFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory: ", err)
		return ""
	}
	return path.Join(cwd, "tasks.json")
}

func ReadTasksFromFile() ([]Task, error) {
	fmt.Println("Starting ReadTasksFromFile")
	filePath := tasksFilePath()
	fmt.Println("File path:", filePath)
	file, err := os.Open(filePath)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist, creating file and returning empty task list")
			file, err = os.Create(filePath)
			os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())
			if err != nil {
				fmt.Println("Error creating file:", err)
				return nil, fmt.Errorf("error creating file: %w", err)
			}
			defer file.Close()
			return []Task{}, nil
		}
		fmt.Println("Error opening file:", err)
		return nil, fmt.Errorf("error opening file: %w", err)
	}
	defer file.Close()
	fmt.Println("File opened successfully")

	tasks := []Task{}
	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&tasks); err != nil {
		fmt.Println("Error decoding tasks from JSON:", err)
		return nil, fmt.Errorf("error decoding tasks from JSON: %w", err)
	}
	fmt.Println("Tasks decoded successfully")

	return tasks, nil
}

func WriteTasksToFile(tasks []Task) error {
	filePath := tasksFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("error creating file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(tasks); err != nil {
		return fmt.Errorf("error encoding tasks to JSON: %w", err)
	}

	return nil
}
