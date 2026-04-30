package cmd

import (
	"fmt"

	"github.com/beeploop/footick/internal/model"
	"github.com/beeploop/footick/internal/utils"
	"github.com/spf13/cobra"
)

var summaryFilter model.SummaryFilter

var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Show total tracked time for a given period",
	Long: `Displays a summary of tracked work durations grouped by task over a selected time range.

This command is intended for billing cycles, invoice preparation, and reviewing productivity.

Supported ranges may include today, week, month, or custom date intervals.

The output should include both per-task totals and an overall total duration, ideally formatted in decimal hours for easy invoicing.

Examples:
chronobar summary today
chronobar summary month`,
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

		utils.PrintJSON(result)
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
