package checkinandout

import "sirclo/entities"

type CheckinAndOut interface {
	Checkin(attendanceId, userId int, temperature float64, status string) error
	Checkout(attendanceId, userId int) error
	GetAllCheck(offset int) ([]entities.GetAllCheck, error)
	GetCheckbyId(id int) (entities.CheckinAndOutResponseFormat, error)
	GetCheckDate(id int) error
	GetTotalPage() (int, error)
}
