package entities

// dipisah folder pisah file juga
type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	ImageUrl string `json:"image_url" form:"image_url"`
}

type UserResponseFormat struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	ImageUrl string `json:"image_url" form:"image_url"`
}
