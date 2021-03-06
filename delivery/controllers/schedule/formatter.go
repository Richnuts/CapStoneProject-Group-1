package schedule

import "time"

type ScheduleRequestFormat struct {
	OfficeId      int        `json:"office_id" form:"office_id"`
	TotalCapacity int        `json:"total_capacity" form:"total_capacity"`
	Month         time.Month `json:"month" form:"month"`
	Year          int        `json:"year" form:"year"`
}

type ScheduleEditFormat struct {
	TotalCapacity int `json:"total_capacity" form:"total_capacity"`
}
