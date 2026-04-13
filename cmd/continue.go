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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// continueCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// continueCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
