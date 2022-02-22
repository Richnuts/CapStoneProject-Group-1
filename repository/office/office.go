package office

import (
	"database/sql"
	"sirclo/entities"
)

type OfficeRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *OfficeRepository {
	return &OfficeRepository{db: db}
}

func (or OfficeRepository) GetOffices() ([]entities.Office, error) {
	var offices []entities.Office
	result, err_offices := or.db.Query(`
	SELECT
		id, name, description
	FROM
		offices`)
	if err_offices != nil {
		return offices, err_offices
	}
	defer result.Close()
	for result.Next() {
		var office entities.Office
		err := result.Scan(&office.Id, &office.Name, &office.Detail)
		if err != nil {
			return offices, err
		}
		offices = append(offices, office)
	}
	return offices, nil
}

func (or OfficeRepository) GetOffice(officeId int) (entities.Office, error) {
	var office entities.Office
	result := or.db.QueryRow("SELECT id, name, description FROM offices WHERE id = ?", officeId)
	err := result.Scan(&office.Id, &office.Name, &office.Detail)
	if err != nil {
		return office, err
	}
	return office, nil
}
