package entities

type Attendance struct {
	Id               int `json:"id" form:"id"`
	Schedule         Schedule
	User             User
	Admin            User
	ImageUrl         string  `json:"image_url" form:"image_url"`
	Description      int     `json:"description" form:"description"`
	Status           int     `json:"status" form:"status"`
	StatusInfo       int     `json:"statusinfo" form:"statusinfo"`
	Checkin          string  `json:"checkin" form:"checkin"`
	Checkout         string  `json:"checkout" form:"checkout"`
	CheckTemperature float64 `json:"checktemperature" form:"checktemperature"`
	CheckStatus      string  `json:"checkstatus" form:"checkstatus"`
}

type CheckinAndOutResponseFormat struct {
	Checkin          *string  `json:"checkin" form:"checkin"`
	Checkout         *string  `json:"checkout" form:"checkout"`
	CheckTemperature *float64 `json:"checktemperature" form:"checktemperature"`
	CheckStatus      *string  `json:"checkstatus" form:"checkstatus"`
}
