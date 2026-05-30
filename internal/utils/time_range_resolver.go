package utils

import (
	"fmt"
	"strings"
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
	if filter.Since != "" {
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
		if !IsValidDateString(filter.From) || !IsValidDateString(filter.To) {
			return model.TimeRange{}, fmt.Errorf("invalid date format provided")
		}

		from, err := time.Parse("2006-01-02", filter.From)
		if err != nil {
			return model.TimeRange{}, err
		}

		to, err := time.Parse("2006-01-02", filter.To)
		if err != nil {
			return model.TimeRange{}, err
		}

		return model.TimeRange{From: from, To: to}, nil

	case filter.Since != "":
		var from time.Time
		var err error

		now := time.Now()

		if IsValidDateString(filter.Since) {
			from, err = time.Parse("2006-01-02", filter.Since)
			if err != nil {
				return model.TimeRange{}, err
			}
		} else {
			from, err = LastWeekdayBeforeN(filter.Since, now)
			if err != nil {
				return model.TimeRange{}, err
			}
		}

		return model.TimeRange{From: from, To: now}, nil

	default:
		return model.TimeRange{}, fmt.Errorf("invalid time range filter")
	}
}

func LastWeekdayBeforeN(weekday string, now time.Time) (time.Time, error) {
	weekdayMap := map[string]time.Weekday{
		"sunday":    time.Sunday,
		"monday":    time.Monday,
		"tuesday":   time.Tuesday,
		"wednesday": time.Wednesday,
		"thursday":  time.Thursday,
		"friday":    time.Friday,
		"saturday":  time.Saturday,
	}

	weekday = strings.ToLower(strings.TrimSpace(weekday))

	weeksBack := 0
	for strings.HasPrefix(weekday, "last-") {
		weeksBack++
		weekday = strings.TrimPrefix(weekday, "last-")
	}

	target, ok := weekdayMap[weekday]
	if !ok {
		return time.Time{}, fmt.Errorf("invalid weekday: %s", weekday)
	}

	today := now.Weekday()
	diff := int(today - target)
	if diff <= 0 {
		diff += 7
	}

	diff += 7 * weeksBack

	return now.AddDate(0, 0, -diff), nil
}
