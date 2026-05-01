package cmd

import (
	"fmt"
	"os"

	"github.com/beeploop/trackbar/internal/app"
	"github.com/spf13/cobra"
)

var application *app.App

var rootCmd = &cobra.Command{
	Use:   "trackbar",
	Short: "Simple freelance time tracking from your terminal",
	Long: `Trackbar is a lightweight CLI time tracker designed for freelancers and developers who want fast, local, and interruption-friendly task tracking.

It supports starting, pausing, resuming, and stopping task timers.

Summaries are optimized for billing workflows, making it easy to convert tracked hours into invoice-ready reports.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		application, err = app.Bootstrap()
		return err
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.Long)
	},
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
