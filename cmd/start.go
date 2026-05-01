package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start tracking time for a new task",
	Long: `Starts a new task timer and begins recording time immediately.

If another task is currently active, the command should prevent starting a new one until the active task is paused or stopped.

The task description is stored in the local SQLite database and a new session entry is created for accurate time tracking.

Use this command whenever you begin work on a new task.

Example:
trackbar start "Description of your task"`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		description := args[0]
		task, err := application.Tracker.CreateTask(description)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Started task #%d: %s\n", task.ID, task.Description)
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
