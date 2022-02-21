package auth

import "sirclo/entities"

type Auth interface {
	Login(email string) (entities.Login, error)
	// Register(user entities.User) error
	FindUserByEmail(email string) (entities.User, error)
}
