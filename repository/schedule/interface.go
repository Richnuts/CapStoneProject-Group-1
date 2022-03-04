package schedule

import (
	"sirclo/entities"
	"time"
)

type Schedule interface {
	CreateSchedule(month time.Month, year int, capacity int, officeId int) error
	EditSchedule(scheduleId int, capacity int) error
	GetSchedule(scheduleId int, offset int) (entities.ScheduleResponse, error)
	GetTotalData(scheduleId int) (int, error)
	GetSchedulesByMonthAndYear(month int, year int, officeId int) ([]entities.Schedule, error)
}
