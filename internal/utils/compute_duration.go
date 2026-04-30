package utils

import (
	"fmt"
	"time"
)

func ComputeDuration(start, end time.Time) (float64, error) {
	if start.IsZero() {
		return 0, fmt.Errorf("invalid start time")
	}

	if end.IsZero() {
		return 0, nil
	}

	return end.Sub(start).Seconds(), nil
}
