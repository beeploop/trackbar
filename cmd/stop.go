package cmd

import (
	"fmt"

	"github.com/beeploop/trackbar/internal/model"
	"github.com/beeploop/trackbar/internal/service"
	"github.com/beeploop/trackbar/internal/utils"
	"github.com/spf13/cobra"
)

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the active task and finalize tracked time",
	Long: `Stops the active task timer, closes the current session, and marks the task as completed.

Once stopped, the task’s total tracked time becomes available in list and summary reports.

Use this command when the task is fully done and no additional work time should be recorded.

Example:
trackbar stop`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		task, err := application.Tracker.StopTask()
		if err != nil {
			fmt.Println(err)
			return
		}

		timerange, err := utils.TimeRangeResolver(&model.SummaryFilter{
			TaskID: task.ID,
			Today:  true,
		})
		if err != nil {
			fmt.Println(err)
			return
		}

		result, err := application.Tracker.SummarizeTask(task.ID, timerange, true)
		if err != nil {
			fmt.Println(err)
			return
		}

		service.NewPrinter().PrintSummary(result)
	},
}

func init() {
	rootCmd.AddCommand(stopCmd)
}
