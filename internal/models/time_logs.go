package models

import "time"

type TimeLog struct {
	ID        int
	TaskID    int
	StartTime time.Time
	EndTime   time.Time
}
