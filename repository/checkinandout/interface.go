package checkinandout

import (
	"sirclo/entities"
)

type CheckinAndOut interface {
	Checkin(check entities.CheckinAndOutResponseFormat) error
	Checkout(check entities.CheckinAndOutResponseFormat) error
}
