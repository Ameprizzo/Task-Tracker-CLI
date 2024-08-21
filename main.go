package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "task-cli",
		Short: "task-cli is a CLI task tracker",
	}

	// Add Task CLI command
	var addsCmd = &cobra.Command{
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
	rootCmd.AddCommand(addsCmd)
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
