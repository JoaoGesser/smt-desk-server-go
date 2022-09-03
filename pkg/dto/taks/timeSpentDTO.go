package taks

import "time"

type TimeSpentDTO struct {
	DateTimeStart  time.Time
	DateTimeFinish time.Time
	TotalSpent     time.Time
}
