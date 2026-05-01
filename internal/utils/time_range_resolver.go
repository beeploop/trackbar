package utils

import (
	"fmt"
	"time"

	"github.com/beeploop/trackbar/internal/model"
)

func TimeRangeResolver(filter *model.SummaryFilter) (model.TimeRange, error) {
	now := time.Now()

	flagCounter := 0
	if filter.Today {
		flagCounter++
	}
	if filter.From != "" || filter.To != "" {
		flagCounter++
	}

	if flagCounter > 1 {
		return model.TimeRange{}, fmt.Errorf("only one time filter can be used at a time")
	}

	switch {
	case filter.Today:
		start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		return model.TimeRange{From: start, To: now}, nil

	case filter.From != "" && filter.To != "":
		from, err := time.Parse("2006-01-02", filter.From)
		if err != nil {
			return model.TimeRange{}, err
		}

		to, err := time.Parse("2006-01-02", filter.To)
		if err != nil {
			return model.TimeRange{}, err
		}

		return model.TimeRange{From: from, To: to}, nil

	default:
		return model.TimeRange{}, fmt.Errorf("invalid time range filter")
	}
}
