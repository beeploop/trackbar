package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the active task and finalize tracked time",
	Long: `Stops the active task timer, closes the current session, and marks the task as completed.

Once stopped, the task’s total tracked time becomes available in list and summary reports.

Use this command when the task is fully done and no additional work time should be recorded.

Example:
footick stop`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("stop called")
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
