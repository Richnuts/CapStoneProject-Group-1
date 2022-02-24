package entities

type Attendance struct {
	Id          int    `json:"id" form:"id"`
	ScheduleId  int    `json:"schedule" form:"schedule"`
	UserId      int    `json:"user" form:"user"`
	AdminId     int    `json:"admin" form:"admin"`
	ImageUrl    string `json:"image_url" form:"image_url"`
	Description int    `json:"description" form:"description"`
	Status      int    `json:"status" form:"status"`
	StatusInfo  int    `json:"statusinfo" form:"statusinfo"`
}

type PendingAttendance struct {
	User        User   `json:"user_id" form:"user_id"`
	ImageUrl    string `json:"image_url" form:"image_url"`
	Description string `json:"description" form:"description"`
}
