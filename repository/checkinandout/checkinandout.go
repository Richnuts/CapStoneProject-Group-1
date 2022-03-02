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

func (cr *CheckRepository) Checkin(attendanceId, userId int, temperature float64, status string) error {
	result, err := cr.db.Exec("UPDATE attendances SET check_in = now(), check_temperature = ?, check_status= ? WHERE id = ? AND user_id = ? AND status = ? AND check_in is NULL ", temperature, status, attendanceId, userId, "Approved")
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

func (cc *CheckRepository) GetCheckbyId(id int) (entities.CheckinAndOutResponseFormat, error) {
	var check entities.CheckinAndOutResponseFormat
	result, err_check := cc.db.Query(`
	SELECT
		schedules.date, attendances.check_in, attendances.check_temperature, attendances.check_out, attendances.check_status
	FROM
		attendances
	JOIN
		schedules on schedules.id = attendances.schedule_id
	WHERE
		id = ?`, id)
	if err_check != nil {
		return check, err_check
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&check.ScheduleDate, &check.Checkin, &check.CheckTemperature, &check.Checkout, &check.CheckStatus)
		if err != nil {
			return check, err
		}
	}
	return check, nil
}

func (cc *CheckRepository) GetCheckDate(id int) error {
	result, err_check := cc.db.Query(`
	SELECT
		schedules.id
	FROM
		attendances
	JOIN
		schedules ON attendances.schedule_id = schedules.id 
	WHERE
		attendances.id = ? AND date(schedules.date) = current_date()`, id)
	if err_check != nil {
		return err_check
	}
	defer result.Close()
	for result.Next() {
		return fmt.Errorf("udah sehat, jangan aneh-aneh")
	}
	return nil
}
