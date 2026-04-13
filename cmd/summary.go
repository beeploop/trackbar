package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

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
		fmt.Println("summary called")
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)
}
