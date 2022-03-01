package auth

type LoginRequestFormat struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

type UserRequestFormat struct {
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
	Email    string `json:"email" form:"email"`
	Nik      string `json:"nik" form:"nik"`
	OfficeId int    `json:"office_id" form:"office_id"`
}
