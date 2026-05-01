package cmd

import (
	"fmt"

	"github.com/beeploop/trackbar/internal/model"
	"github.com/beeploop/trackbar/internal/service"
	"github.com/beeploop/trackbar/internal/utils"
	"github.com/spf13/cobra"
)

var summaryFilter model.SummaryFilter

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show total tracked time for a given period",
	Long: `Displays a summary of tracked work durations grouped by task over a selected time range.

This command is intended for billing cycles, invoice preparation, and reviewing productivity.

The output should include both per-task totals and an overall total duration, ideally formatted in decimal hours for easy invoicing.

Examples:
trackbar summary --today
trackbar summary --from '2026-04-14' --to '2026-04-27'
trackbar summary --from '2026-04-14' --to '2026-04-27' --completed-only
`,
	Run: func(cmd *cobra.Command, args []string) {
		timerange, err := utils.TimeRangeResolver(&summaryFilter)
		if err != nil {
			fmt.Println(err)
			return
		}

		result, err := application.Tracker.SummarizeTask(summaryFilter.TaskID, timerange, summaryFilter.CompletedOnly)
		if err != nil {
			fmt.Println(err)
			return
		}

		service.NewPrinter().PrintSummary(result)
	},
}

func init() {
	// task filters
	summaryCmd.Flags().IntVarP(&summaryFilter.TaskID, "task", "t", 0, "Specify task ID")

	// time filters
	summaryCmd.Flags().BoolVar(&summaryFilter.Today, "today", false, "Show today's summary")
	summaryCmd.Flags().StringVar(&summaryFilter.From, "from", "", "Specify start date (YYYY-MM-DD)")
	summaryCmd.Flags().StringVar(&summaryFilter.To, "to", "", "Specify end date (YYYY-MM-DD)")
	summaryCmd.Flags().BoolVar(&summaryFilter.CompletedOnly, "completed-only", false, "Only include tasks that are marked completed")

	rootCmd.AddCommand(summaryCmd)
}
