package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var pauseCmd = &cobra.Command{
	Use:   "pause",
	Short: "Pause the active task timer",
	Long: `Pauses the currently active task by ending its active session while keeping the task open for later continuation.

This command is useful when switching context, taking breaks, or temporarily stopping work without marking the task as complete.

The elapsed time from the current session is saved automatically.

Example:
trackbar pause`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		task, err := application.Tracker.PauseTask()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Paused tracking for the task with ID: %d\n", task.ID)
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)
}
