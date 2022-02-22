package checkinandout

type CheckinAndOut interface {
	Checkin(attendanceId, userId int, temperature float64, status string) error
	Checkout(attendanceId, userId int) error
}
