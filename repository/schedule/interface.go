package schedule

import (
	"sirclo/entities"
	"time"
)

type Schedule interface {
	CreateSchedule(month time.Month, year int, capacity int, officeId int) error
	EditSchedule(date string, capacity int, officeId int) error
	GetSchedule(scheduleId int) (entities.ScheduleResponse, error)
}
