package entities

type Certificate struct {
	Id          int    `json:"id" form:"id"`
	UserId      int    `json:"userid" form:"usersid"`
	ImageURL    string `json:"imageurl" form:"imageurl"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}

type CertificateResponse struct {
	Id          int `json:"id" form:"id"`
	User        User
	ImageURL    string `json:"imageurl" form:"imageurl"`
	Status      string `json:"status" form:"status"`
	Description string `json:"description" form:"description"`
}
