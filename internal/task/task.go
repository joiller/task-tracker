package task

import (
	"fmt"
	"time"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_DONE        TaskStatus = "done"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in_progress"
)

type Task struct {
	ID          int64      `json:"id,omitempty"`
	Description string     `json:"description,omitempty"`
	Status      TaskStatus `json:"status,omitempty"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return nil
	}
	filteredTasks := []Task{}
	switch status {
	case "all":
		filteredTasks = tasks
	default:
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}
	fmt.Println()
	fmt.Println("task status: ", status)
	for _, task := range filteredTasks {
		fmt.Printf("taskId: %d, %s, %s\n", task.ID, task.Description, task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
	fmt.Println()
	return nil
}

func statusColor(status TaskStatus) string {
	switch status {
	case TASK_STATUS_TODO:
		return "\033[31m"
	case TASK_STATUS_DONE:
		return "\033[32m"
	case TASK_STATUS_IN_PROGRESS:
		return "\033[33m"
	default:
		return "\033[0m"
	}
}

func TaskStatusFromString(status string) TaskStatus {
	switch status {
	case "todo":
		return TASK_STATUS_TODO
	case "done":
		return TASK_STATUS_DONE
	case "in_progress":
		return TASK_STATUS_IN_PROGRESS
	default:
		return "all"
	}
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	id := int64(1)
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		id = lastTask.ID + 1
	}
	task := NewTask(id, description)
	tasks = append(tasks, *task)
	fmt.Printf("Added task \"%s\" with id %d\n", description, id)
	return WriteTasksToFile(tasks)
}

func UpdateTaskStatus(ids []int64, status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	updatedTasks := []Task{}
	for i, task := range tasks {
		for _, id := range ids {
			if task.ID == id {
				tasks[i].Status = status
				tasks[i].UpdatedAt = time.Now()
				updatedTasks = append(updatedTasks, tasks[i])
				break
			}
		}
	}
	if len(updatedTasks) == 0 {
		return fmt.Errorf("no tasks with the provided ids were found")
	}
	for _, task := range updatedTasks {
		fmt.Printf("Updated taskId: %d, %s, %s\n", task.ID, task.Description, task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
	return WriteTasksToFile(tasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	updatedTasks := []Task{}
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			updatedTasks = append(updatedTasks, tasks[i])
			break
		}
	}
	if len(updatedTasks) == 0 {
		return fmt.Errorf("no tasks with the provided id were found")
	}
	for _, task := range updatedTasks {
		fmt.Printf("Updated taskId: %d, %s, %s\n", task.ID, task.Description, task.UpdatedAt.Format("2006-01-02 15:04:05"))
	}
	return WriteTasksToFile(tasks)

}

func DeleteTasks(ids []int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}
	remainingTasks := []Task{}
	for _, task := range tasks {
		found := false
		for _, id := range ids {
			if task.ID == id {
				found = true
				fmt.Println("find task id: ", task.ID)
				break
			}
		}
		if !found {
			remainingTasks = append(remainingTasks, task)
		}
	}
	if len(remainingTasks) == len(tasks) {
		return fmt.Errorf("no tasks with the provided ids were found")
	}

	return WriteTasksToFile(remainingTasks)

}
