package attendance

type AttendanceRequestFormat struct {
	ScheduleId  int    `json:"schedule_id" form:"schedule_id"`
	Description string `json:"description" form:"description"`
}
