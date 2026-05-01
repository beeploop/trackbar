package cmd

import (
	"fmt"

	"github.com/beeploop/trackbar/internal/service"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List tracked tasks and their current status",
	Long: `Displays tracked tasks grouped by status (paused and active tasks only).

Each task entry includes its description and accumulated tracked duration.
This command provides a quick overview of what you are currently working on.

Example:
trackbar list`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		result, err := application.Tracker.ListTask()
		if err != nil {
			fmt.Println(err)
			return
		}

		service.NewPrinter().PrintTaskList(result)
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
