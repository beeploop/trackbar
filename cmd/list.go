package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tracked tasks and their current status",
	Long: `Displays tracked tasks grouped by their current state, such as active, paused, or recently completed.

Each task entry includes its description and accumulated tracked duration.

This command provides a quick overview of what you are currently working on and what has recently been finished.

Example:
footick list`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("list called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
