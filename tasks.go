package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/charmbracelet/lipgloss"
)

// Task struct represents a task with necessary fields
type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

// GetFormattedTime returns the current time formatted as "2006-01-02 15:04:05"
func GetFormattedTime() string {
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	return formattedTime
}

func statusColor(status string) string {
	switch status {
	case "todo":
		return "5"
	case "in-progress":
		return "202"
	case "done":
		return "#04B575"
	default:
		return "#3C3C3C"
	}
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
	currentTime := GetFormattedTime()
	newTask := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   currentTime,
		UpdatedAt:   currentTime,
	}
	tasks = append(tasks, newTask)
	err = SaveTasks(tasks)
	if err != nil {
		return Task{}, err
	}
	return newTask, nil
}

// UpdateTask updates the description of an existing task
func UpdateTask(id int, description string) (Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return Task{}, err
	}

	// Find the task by ID and update it
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = GetFormattedTime()

			err = SaveTasks(tasks)
			if err != nil {
				return Task{}, err
			}
			return tasks[i], nil
		}
	}
	return Task{}, fmt.Errorf("Task with ID %d not found", id)
}

// DeleteTask deletes a task by its ID
func DeleteTask(id int) error {
	tasks, err := GetTasks()
	if err != nil {
		return err
	}

	//Find the task by its ID and delete it
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return SaveTasks(tasks)
		}
	}
	return fmt.Errorf("Task with ID %d not found", id)
}

// UpdateTaskStatus updates the status of a task and returns the updated task.
func UpdateTaskStatus(id int, newStatus string) (Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return Task{}, err
	}

	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Status = newStatus
			tasks[i].UpdatedAt = GetFormattedTime()

			err = SaveTasks(tasks)
			if err != nil {
				return Task{}, err
			}
			return tasks[i], nil
		}
	}
	return Task{}, fmt.Errorf("Task with ID %d not found", id)
}

// ListTasks lists all tasks, optionally filtered by status
func ListTasks(status string) ([]Task, error) {
	tasks, err := GetTasks()
	if err != nil {
		return nil, err
	}

	// If status is provided, filter tasks by status
	if status != "" {
		var filteredTasks []Task
		for _, task := range tasks {
			if task.Status == status {
				filteredTasks = append(filteredTasks, task)
			}
		}
		fmt.Println()
		fmt.Println(
			lipgloss.NewStyle().
				Bold(true).
				Foreground(lipgloss.Color("#FFCC66")).
				MarginBottom(1).
				Render(fmt.Sprintf("Tasks (%s)", status)))
		return filteredTasks, nil
	}

	return tasks, nil
}
