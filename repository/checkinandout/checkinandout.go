package checkinandout

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type CheckRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CheckRepository {
	return &CheckRepository{db: db}
}

func (cr *CheckRepository) Checkin(attendance entities.CheckinAndOutResponseFormat) error {
	result, err := cr.db.Exec("UPDATE attendances SET check_in = ?, check_temperature = ?, updated_at = now() WHERE id = ? AND deleted_at IS null", attendance.Checkin, attendance.CheckTemperature, attendance.User.Id)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
