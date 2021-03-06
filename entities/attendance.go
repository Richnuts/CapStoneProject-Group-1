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

type CheckinAndOutResponseFormat struct {
	Id               int      `json:"id" form:"id"`
	ScheduleDate     string   `json:"scheduledate" form:"scheduledate"`
	Checkin          *string  `json:"checkin" form:"checkin"`
	Checkout         *string  `json:"checkout" form:"checkout"`
	CheckTemperature *float64 `json:"checktemperature" form:"checktemperature"`
	CheckStatus      *string  `json:"checkstatus" form:"checkstatus"`
}

type GetAllCheck struct {
	Id            int    `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	Nik           string `json:"nik" form:"nik"`
	VaccineStatus string `json:"vaccinestatus" form:"vaccinestatus"`
	CheckData     []CheckinAndOutResponseFormat
}

type GetAllCheckWithPage struct {
	AllCheck  []GetAllCheck
	TotalPage int `json:"totalpage" form:"totalpage"`
}

type GetCheckbyDate struct {
	Id            int    `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	Nik           string `json:"nik" form:"nik"`
	VaccineStatus string `json:"vaccinestatus" form:"vaccinestatus"`
	CheckData     CheckinAndOutResponseFormat
}

type GetCheckbyDateWithPage struct {
	UsersCheck []GetCheckbyDate
	TotalPage  int `json:"totalpage" form:"totalpage"`
}

type PendingAttendance struct {
	Id             int                `json:"id" form:"id"`
	ScheduleId     int                `json:"schedule_id" form:"schedule_id"`
	ActualCapacity int                `json:"actual_capacity" form:"actual_capacity"`
	Date           string             `json:"date" form:"date"`
	Office         string             `json:"office" form:"office"`
	ImageUrl       string             `json:"image_url" form:"image_url"`
	Description    string             `json:"description" form:"description"`
	RequestTime    string             `json:"request_time" form:"request_time"`
	User           UserResponseFormat `json:"user" form:"user"`
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
	TotalData  int                   `json:"total_data" form:"total_data"`
	Attendance []AttendanceGetFormat `json:"attendance" form:"attendance"`
}
type PendingAttendancePageFormat struct {
	TotalPage  int                 `json:"total_page" form:"total_page"`
	TotalData  int                 `json:"total_data" form:"total_data"`
	Attendance []PendingAttendance `json:"attendance" form:"attendance"`
}
