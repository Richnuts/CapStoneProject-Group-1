package entities

type Schedule struct {
	Id            int `json:"id" form:"id"`
	Office        Office
	TotalCapacity int    `json:"totalcapacity" form:"totalcapacity"`
	Capacity      int    `json:"capacity" form:"capacity"`
	Date          string `json:"date" form:"date"`
}
