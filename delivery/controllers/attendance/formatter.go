package attendance

type AttendanceRequestFormat struct {
	ScheduleId  int    `json:"schedule_id" form:"schedule_id"`
	Description string `json:"description" form:"description"`
}

type AttendanceEditFormat struct {
	ScheduleId int    `json:"schedule_id" form:"schedule_id"`
	Status     string `json:"status" form:"status"`
	StatusInfo string `json:"status_info" form:"status_info"`
}


