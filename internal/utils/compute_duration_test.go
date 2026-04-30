package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestComputeDuration(t *testing.T) {
	t.Run("test compute_duration", func(t *testing.T) {
		cases := []struct {
			Start       time.Time
			End         time.Time
			Expected    float64
			ExpectError bool
		}{
			{
				Start:       time.Date(2026, 04, 30, 13, 0, 0, 0, time.Local),
				End:         time.Date(2026, 04, 30, 14, 0, 0, 0, time.Local),
				Expected:    3600,
				ExpectError: false,
			},
			{
				Start:       time.Date(2026, 04, 30, 13, 0, 0, 0, time.Local),
				End:         time.Date(2026, 04, 30, 14, 55, 0, 0, time.Local),
				Expected:    6900,
				ExpectError: false,
			},
			{
				Start:       time.Time{},
				End:         time.Date(2026, 04, 30, 14, 55, 0, 0, time.Local),
				Expected:    0,
				ExpectError: true,
			},
			{
				Start:       time.Date(2026, 04, 30, 14, 55, 0, 0, time.Local),
				End:         time.Time{},
				Expected:    0,
				ExpectError: false,
			},
		}

		for _, tc := range cases {
			duration, err := ComputeDuration(tc.Start, tc.End)
			if tc.ExpectError {
				assert.Error(t, err)
			}
			assert.Equal(t, tc.Expected, duration)
		}
	})
}
