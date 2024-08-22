package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "task-cli",
		Short: "task-cli is a CLI task tracker",
	}

	// Add Task CLI command
	var AddTaskCmd = &cobra.Command{
		Use:   "add [description]",
		Short: "Add new task",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			description := args[0]
			task, err := AddTask(description)
			if err != nil {
				fmt.Println("Error adding task:", err)
				return
			}
			fmt.Printf("Task added successfully(ID: %d)\n", task.ID)
		},
	}

	// Update Task Command
	var updateTaskCmd = &cobra.Command{
		Use:   "update [id] [description]",
		Short: "Update an existing task's description",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID: ", err)
				return
			}
			description := args[1]
			task, err := UpdateTask(id, description)
			if err != nil {
				fmt.Println("Error updating task: ", err)
				return
			}
			fmt.Printf("Task updated successfully (ID: %d)\n", task.ID)
		},
	}

	// Delete Task Command
	var deleteTaskCmd = &cobra.Command{
		Use:   "delete [id]",
		Short: "Delete a task by its ID",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID: ", err)
				return
			}
			err = DeleteTask(id)
			if err != nil {
				fmt.Println("Error deleting task: ", err)
				return
			}
			fmt.Printf("Task deleted successfully (ID: %d)\n", id)
		},
	}

	// Mark Task As In-Progress Command
	var MarkTaskAsInProgressCmd = &cobra.Command{
		Use:   "mark-in-progress [id]",
		Short: "Mark a task as in-progress",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID:", err)
				return
			}
			task, err := UpdateTaskStatus(id, "in-progress")
			if err != nil {
				fmt.Println("Error marking task as in-progress:", err)
				return
			}
			fmt.Printf("Task marked as in-progress (ID: %d)\n", task.ID)
		},
	}

	// Mark Task As Done Command
	var MarkTaskAsDoneCmd = &cobra.Command{
		Use:   "mark-done [id]",
		Short: "Mark a task as done",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			id, err := strconv.Atoi(args[0])
			if err != nil {
				fmt.Println("Invalid task ID:", err)
				return
			}
			task, err := UpdateTaskStatus(id, "done")
			if err != nil {
				fmt.Println("Error marking task as done:", err)
				return
			}
			fmt.Printf("Task marked as done (ID: %d)\n", task.ID)
		},
	}

	// List Tasks Command
	var ListTasksCmd = &cobra.Command{
		Use:   "list [status]",
		Short: "List all tasks, optionally filtered by status (todo, in-progress, done)",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			status := ""
			if len(args) > 0 {
				status = args[0]
			}
			tasks, err := ListTasks(status)
			if err != nil {
				fmt.Println("Error listing tasks:", err)
				return
			}
			if len(tasks) == 0 {
				fmt.Println(
					lipgloss.NewStyle().
						Bold(true).
						Padding(1, 0).
						Foreground(lipgloss.Color("#FFCC66")).
						Render("No tasks found"))
				return
			}
			for _, task := range tasks {
				fmt.Printf("ID: %d, Description: %s, Status: %s, CreatedAt: %s, UpdatedAt: %s\n",
					task.ID, task.Description, task.Status, task.CreatedAt, task.UpdatedAt)
			}
		},
	}

	rootCmd.AddCommand(AddTaskCmd, updateTaskCmd, deleteTaskCmd, MarkTaskAsInProgressCmd, MarkTaskAsDoneCmd, ListTasksCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
