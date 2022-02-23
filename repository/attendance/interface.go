package attendance

type Attendance interface {
	CreateAttendance(userId int, scheduleId int, description string, imageUrl string) error
	GetUserAttendanceStatus(userId int, scheduleId int) error
}
