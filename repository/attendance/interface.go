package attendance

import "sirclo/entities"

type Attendance interface {
	CreateAttendance(userId int, scheduleId int, description string, imageUrl string) error
	EditAttendance(attendanceId int, adminId int, status string, statusInfo string) error
	GetPendingAttendance(offset int) ([]entities.PendingAttendance, error) //sort by earliest created_at
	GetPendingAttendanceTotalPage() (int, error)
	GetMyAttendance(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error)
	GetAttendanceById(attendanceId int) (entities.AttendanceGetFormat, error)
	GetMyAttendanceSortByLatest(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error)  // created at
	GetMyAttendanceSortByLongest(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error) // date
	GetMyAttendanceTotalPage(userId int, status string) (int, error)
	GetUserAttendanceStatus(userId int, scheduleId int) error
	GetUserVaccineStatus(userId int) error
	CheckCapacity(scheduleId int) (int, error)
}
