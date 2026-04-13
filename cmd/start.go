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
footick start "Description of your task"`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("start called")
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
