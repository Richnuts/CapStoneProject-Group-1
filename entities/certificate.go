package entities

type Certificate struct {
	Id          int    `json:"id" form:"id"`
	UserId      int    `json:"userid" form:"usersid"`
	ImageURL    string `json:"imageurl" form:"imageurl"`
	VaccineDose int    `json:"vaccinedose" form:"vaccinedose"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}

type CertificateResponseGetByIdAndUID struct {
	Id          int    `json:"id" form:"id"`
	ImageURL    string `json:"imageurl" form:"imageurl"`
	VaccineDose int    `json:"vaccinedose" form:"vaccinedose"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}

type CertificateResponsePending struct {
	Id          int    `json:"id" form:"id"`
	ImageURL    string `json:"imageurl" form:"imageurl"`
	VaccineDose int    `json:"vaccinedose" form:"vaccinedose"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}