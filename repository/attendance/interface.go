package attendance

import (
	"sirclo/entities"
	"time"
)

type Attendance interface {
	CreateAttendance(userId int, scheduleId int, description string, imageUrl string) error
	EditAttendance(attendanceId int, adminId int, status string, statusInfo string) error
	GetPendingAttendance(offset int, officeId int) ([]entities.PendingAttendance, error) //sort by earliest created_at
	GetPendingAttendanceTotalData(officeId int) (int, error)
	GetMyAttendance(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error)
	GetAttendanceById(attendanceId int) (entities.AttendanceGetFormat, error)
	GetMyAttendanceSortByLatest(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error)  // created at
	GetMyAttendanceSortByLongest(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error) // date
	GetMyAttendanceTotalData(userId int, status string) (int, error)
	GetUserAttendanceStatus(userId int, scheduleId int) error
	GetUserVaccineStatus(userId int) error
	CheckCapacity(scheduleId int) (int, error)
	CheckCreateRequestDate(scheduleId int) (time.Time, error)
}
