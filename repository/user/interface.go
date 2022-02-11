package user

import (
	"sirclo/entities"
)

type User interface {
	GetUser(id int) (entities.UserResponseFormat, error)
	DeleteUser(id int) error
	EditUser(user entities.User, id int) error
}
