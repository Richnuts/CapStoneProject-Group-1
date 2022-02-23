package checkinandout

type CheckinRequestFormat struct {
	Id          int     `json:"id" form:"id"`
	Temperature float64 `json:"temperature" form:"temperature"`
}

type CheckoutRequestFormat struct {
	Id int `json:"id" form:"id"`
}
