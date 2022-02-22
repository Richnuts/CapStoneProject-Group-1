package entities

// dipisah folder pisah file juga
type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	Nik      string `json:"nik" form:"nik"`
	Role     string `json:"role" form:"role"`
	ImageUrl string `json:"image_url" form:"image_url"`
	OfficeId int    `json:"office_id" form:"office_id"`
}

type UserResponseFormat struct {
	Id            int    `json:"id" form:"id"`
	Name          string `json:"name" form:"name"`
	Email         string `json:"email" form:"email"`
	ImageUrl      string `json:"image_url" form:"image_url"`
	Nik           string `json:"nik" form:"nik"`
	VaccineStatus string `json:"vaccine_status" form:"vaccine_status"`
	Office        string `json:"office" form:"office"`
}
