package service

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/beeploop/trackbar/internal/model"
	"github.com/beeploop/trackbar/internal/utils"
)

type Printer struct {
	Writer *tabwriter.Writer
}

func NewPrinter() *Printer {
	return &Printer{
		Writer: tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0),
	}
}

func (p *Printer) PrintTaskList(tasks []model.TaskSession) {
	paused := make([]model.TaskSession, 0)
	active := make([]model.TaskSession, 0)

	for _, task := range tasks {
		switch task.Task.Status {
		case model.TASK_PAUSED:
			paused = append(paused, task)
		case model.TASK_ACTIVE:
			active = append(active, task)
		}
	}

	p.printGroup("PAUSED", paused)
	p.printGroup("ACTIVE", active)
}

func (p *Printer) PrintSummary(tasks []model.TaskSession) {
	fmt.Fprintln(p.Writer, "\nID\tDESCRIPTION\tSESSIONS\tSTATUS\tTOTAL")

	totalDuration := 0.0
	for _, task := range tasks {
		taskDuration := 0.0

		for _, session := range task.Sessions {
			duration, err := utils.ComputeDuration(session.StartedAt, session.EndedAt)
			if err != nil {
				continue
			}
			taskDuration += duration
		}
		totalDuration += taskDuration

		fmt.Fprintf(
			p.Writer,
			"#%d\t%s\t(%d) sessions\t%s\t%s\n",
			task.Task.ID,
			task.Task.Description,
			len(task.Sessions),
			task.Task.Status,
			utils.FormatHMS(taskDuration),
		)
	}

	fmt.Fprintf(p.Writer, "\nTOTAL:\t%s\n", utils.FormatHMS(totalDuration))
	p.Writer.Flush()
}

func (p *Printer) printGroup(groupName string, slice []model.TaskSession) {
	fmt.Printf("\n%s\n", groupName)
	fmt.Fprintln(p.Writer, "ID\tDESCRIPTION\tSTART\tEND\tTOTAL")
	for _, ts := range slice {
		for i, session := range ts.Sessions {
			var id, description string

			if i == 0 {
				id = fmt.Sprintf("#%d", ts.Task.ID)
				description = ts.Task.Description
			} else {
				id = ""
				description = ""
			}

			duration, err := utils.ComputeDuration(session.StartedAt, session.EndedAt)
			if err != nil {
				continue
			}

			fmt.Fprintf(
				p.Writer,
				"%s\t%s\t%s\t%s\t%s\n",
				id,
				description,
				utils.FormatTime(session.StartedAt),
				utils.FormatTime(session.EndedAt),
				utils.FormatHMS(duration),
			)
		}
	}
	p.Writer.Flush()
}
