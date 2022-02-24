package attendance

type Attendance interface {
	CreateAttendance(userId int, scheduleId int, description string, imageUrl string) error
	// ApproveAttendance(attendanceId int) error
	// GetPendingAttendance(scheduleId int) []entities.PendingAttendance //sort by earliest created_at
	// GetMyAttendance(scheduleId int) ([]entities.Attendance, error)
	// GetMyAttendanceById(scheduleId int, attendanceId int) (entities.Attendance, error)
	// GetMyAttendanceSortByLatest(scheduleId int) ([]entities.Attendance, error)  // created at
	// GetMyAttendanceSortByLongest(scheduleId int) ([]entities.Attendance, error) // date
	GetUserAttendanceStatus(userId int, scheduleId int) error
	GetUserVaccineStatus(userId int) error
	// GetCheckDataById(attendanceId, scheduleId int) (entities.Attendance, error)
}
