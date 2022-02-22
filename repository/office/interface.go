package office

import "sirclo/entities"

type Office interface {
	GetOffices() ([]entities.Office, error)
	GetOffice(officeId int) (entities.Office, error)
}
