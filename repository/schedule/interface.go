package schedule

import "time"

type Schedule interface {
	CreateSchedule(month time.Month, year int, capacity int, officeId int) error
	EditSchedule(date string, capacity int, officeId int) error
}
