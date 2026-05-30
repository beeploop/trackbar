package utils

import "time"

func IsValidDateString(str string) bool {
	_, err := time.Parse("2006-01-02", str)
	if err != nil {
		return false
	}
	return true
}
