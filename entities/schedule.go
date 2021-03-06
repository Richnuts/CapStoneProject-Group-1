package entities

import "time"

type Schedule struct {
	Id            int       `json:"id" form:"id"`
	TotalCapacity int       `json:"total_capacity" form:"total_capacity"`
	Capacity      int       `json:"capacity" form:"capacity"`
	OfficeId      int       `json:"office" form:"office"`
	Date          time.Time `json:"time" form:"time"`
}

type ScheduleResponse struct {
	Id            int                  `json:"id" form:"id"`
	Date          time.Time            `json:"date" form:"date"`
	TotalCapacity int                  `json:"total_capacity" form:"total_capacity"`
	Capacity      int                  `json:"capacity" form:"capacity"`
	Office        string               `json:"office" form:"office"`
	TotalPage     int                  `json:"total_page" form:"total_page"`
	TotalData     int                  `json:"total_data" form:"total_data"`
	Attendance    []UserResponseFormat `json:"user" form:"user"`
}
