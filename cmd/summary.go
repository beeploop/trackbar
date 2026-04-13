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

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// summaryCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// summaryCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
