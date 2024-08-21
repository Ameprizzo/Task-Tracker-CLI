package main

import (
	"encoding/json"
	"os"
	"time"
)

// Task struct represents a task with necessary fields
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// Get all tasks from the JSON file
func GetTasks() ([]Task, error) {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}, nil // No tasks.json file, return an empty list
		}
		return nil, err
	}
	var tasks []Task
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		return nil, err
	}
	return tasks, nil

}

// SaveTasks saves tasks to the JSON file
func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		return err
	}
	return nil

}

// AddTask adds a new task to the task list
func AddTask(description string) (Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return Task{}, err
	}

	id := len(tasks) + 1
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	newTask := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   formattedTime,
		UpdatedAt:   formattedTime,
	}
	tasks = append(tasks, newTask)
	err = SaveTasks(tasks)
	if err != nil {
		return Task{}, err
	}
	return newTask, nil
}
