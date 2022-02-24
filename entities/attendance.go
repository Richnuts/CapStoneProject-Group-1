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
	Id          int                `json:"id" form:"id"`
	Date        string             `json:"date" form:"date"`
	ImageUrl    string             `json:"image_url" form:"image_url"`
	Description string             `json:"description" form:"description"`
	RequestTime string             `json:"request_time" form:"request_time"`
	User        UserResponseFormat `json:"user" form:"user"`
}

type AttendanceGetFormat struct {
	Id           string  `json:"id" form:"id"`
	Name         string  `json:"name" form:"name"`
	Date         string  `json:"date" form:"date"`
	Office       string  `json:"office" form:"office"`
	Status       string  `json:"status" form:"status"`
	StatusInfo   string  `json:"status_info" form:"status_info"`
	RequestTime  string  `json:"request_time" form:"request_time"`
	ApprovedTime *string `json:"approved_time" form:"approved_time"`
	AdminName    *string `json:"admin_name" form:"admin_name"`
	CheckIn      *string `json:"check_in" form:"check_in"`
}

type AttendancePageFormat struct {
	TotalPage  int                   `json:"total_page" form:"total_page"`
	Attendance []AttendanceGetFormat `json:"attendance" form:"attendance"`
}
type PendingAttendancePageFormat struct {
	TotalPage  int                 `json:"total_page" form:"total_page"`
	Attendance []PendingAttendance `json:"attendance" form:"attendance"`
}
