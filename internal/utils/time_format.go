package utils

import (
	"fmt"
	"time"
)

func FormatTime(t time.Time) string {
	if t.IsZero() {
		return "---"
	}

	return t.Format("January 2, 2006 3:04 PM")
}

func FormatHMS(seconds float64) string {
	duration := time.Duration(seconds) * time.Second

	h := int(duration.Hours())
	m := int(duration.Minutes()) % 60
	s := int(duration.Seconds()) % 60

	return fmt.Sprintf("%dh %dm %ds", h, m, s)
}
