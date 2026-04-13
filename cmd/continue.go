package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var continueCmd = &cobra.Command{
	Use:   "continue",
	Short: "Resume the most recently paused task",
	Long: `Resumes tracking for the most recently paused task by creating a new session segment linked to the same task.

This allows a single task to accumulate time across multiple work intervals while preserving accurate timing history.

If no paused task exists, the command should return a helpful message.

Example:
footick continue`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("continue called")
	},
}

func init() {
	rootCmd.AddCommand(continueCmd)
}
