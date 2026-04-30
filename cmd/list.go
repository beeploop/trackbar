package cmd

import (
	"fmt"

	"github.com/beeploop/footick/internal/utils"
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
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := application.Tracker.ListTask()
		if err != nil {
			fmt.Println(err)
			return
		}

		utils.PrintJSON(tasks)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
