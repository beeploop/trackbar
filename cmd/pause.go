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
footick pause`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("pause called")
	},
}

func init() {
	rootCmd.AddCommand(pauseCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// pauseCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// pauseCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
