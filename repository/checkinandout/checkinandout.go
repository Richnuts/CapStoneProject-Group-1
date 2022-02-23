package checkinandout

import (
	"database/sql"
	"fmt"
)

type CheckRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CheckRepository {
	return &CheckRepository{db: db}
}

func (cr *CheckRepository) Checkin(attendanceId, userId int, temperature float64, status string) error {
	result, err := cr.db.Exec("UPDATE attendances SET check_in = now(), check_temperature = ?, check_status= ? WHERE id = ? AND user_id = ? AND status = ? AND check_in is NULL ", temperature, status, attendanceId, userId, "Approved")
	fmt.Println(temperature, status, attendanceId, userId)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("wfo request not found / already checked in")
	}
	return nil
}

func (cr *CheckRepository) Checkout(attendanceId, userId int) error {
	result, err := cr.db.Exec("UPDATE attendances SET check_out = now() WHERE id = ? AND user_id = ? AND status = ? AND check_in is NOT NULL AND check_out is NULL", attendanceId, userId, "Approved")
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("wfo request not found")
	}
	return nil
}
