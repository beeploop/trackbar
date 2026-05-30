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

The --since flag is a shortcut for:
	--from <DATE> --to <NOW>

	Accepted values:
		2026-04-14
		monday
		friday
		last-monday
		last-last-monday

	Each 'last-' prefix moves the reference back one week.

Examples:
trackbar summary --today
trackbar summary --from '2026-04-14' --to '2026-04-27'
trackbar summary --from '2026-04-14' --to '2026-04-27' --completed-only
trackbar summary --since '2026-04-14'
trackbar summary --since monday
trackbar summary --since last-monday
trackbar summary --since last-last-monday
`,
	Run: func(cmd *cobra.Command, args []string) {
		timerange, err := utils.TimeRangeResolver(&summaryFilter)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("showing results: %s - %s\n", utils.FormatTime(timerange.From), utils.FormatTime(timerange.To))

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
	summaryCmd.Flags().StringVar(&summaryFilter.Since, "since", "", "Start date or weekday shortcut")
	summaryCmd.Flags().BoolVar(&summaryFilter.CompletedOnly, "completed-only", false, "Only include tasks that are marked completed")

	rootCmd.AddCommand(summaryCmd)
}
