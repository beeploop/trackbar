package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var continueCmd = &cobra.Command{
	Use:   "continue",
	Short: "Resume the most recently paused task",
	Long: `Resumes tracking for the most recently paused task by creating a new session segment linked to the same task.

This allows a single task to accumulate time across multiple work intervals while preserving accurate timing history.

If no paused task exists, the command should return a helpful message.

Example:
trackbar continue [Task ID]`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		taskID, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Println("provided task ID is not valid")
			return
		}

		task, err := application.Tracker.ContinueTask(taskID)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Continued tracking task with ID: %d\n", task.ID)
	},
}

func init() {
	rootCmd.AddCommand(continueCmd)
}
