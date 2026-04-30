package model

type SummaryFilter struct {
	TaskID        int
	Today         bool
	From          string
	To            string
	CompletedOnly bool
}
