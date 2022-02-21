package entities

type Office struct {
	Id     int    `json:"id" form:"id"`
	Name   string `json:"name" form:"name"`
	Detail string `json:"detail" form:"detail"`
}

type OfficeUserResponse struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}
