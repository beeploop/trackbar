package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var switchCmd = &cobra.Command{
	Use:   "switch",
	Short: "Switch to another task by pausing the current one and starting another.",
	Long: `Switch pauses the currently active task and immediately starts a new session for the target task.

If a task is currently active, its running session will be stopped at the time of the switch.
A new session will then be started for the target task.

If no task is active, switch behaves like continue.

If the target task is already active, no changes are made.

This command provides a fast way to move between tasks without manually pausing and continuing.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return fmt.Errorf("switch requires a task ID")
		}

		targetTaskID, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("Invalid task ID: must be a number")
		}

		task, err := application.Tracker.SwitchTask(targetTaskID)
		if err != nil {
			return err
		}

		fmt.Printf("switched task ID: %d\n", task.ID)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(switchCmd)
}
