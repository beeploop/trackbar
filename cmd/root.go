package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "footick",
	Short: "Simple freelance time tracking from your terminal",
	Long: `Footick is a lightweight CLI time tracker designed for freelancers and developers who want fast, local, and interruption-friendly task tracking.

It supports starting, pausing, resuming, and stopping task timers.

Summaries are optimized for billing workflows, making it easy to convert tracked hours into invoice-ready reports.`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
