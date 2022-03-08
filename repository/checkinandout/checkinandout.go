package checkinandout

import (
	"database/sql"
	"fmt"
	"math"
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
		attendances.id, CONVERT_TZ(schedules.date, '+00:00', '+7:00'), attendances.check_in, attendances.check_temperature, attendances.check_out, attendances.check_status
	FROM
		attendances
	JOIN
		schedules on schedules.id = attendances.schedule_id
	WHERE
		attendances.id = ? AND attendances.check_in is not null AND attendances.check_out is not null`, id)
	if err_check != nil {
		return check, err_check
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&check.Id, &check.ScheduleDate, &check.Checkin, &check.CheckTemperature, &check.Checkout, &check.CheckStatus)
		if err != nil {
			return check, err
		}
	}
	return check, nil
}

func (cc *CheckRepository) GetAllCheck(id, offset int) ([]entities.GetCheckbyDate, error) {
	var hasil []entities.GetCheckbyDate
	result, err_check := cc.db.Query(`
	SELECT
		users.id, users.name, users.nik, users.vaccine_status, attendances.id, CONVERT_TZ(schedules.date, '+00:00', '+7:00'), attendances.check_in, attendances.check_temperature, attendances.check_out, attendances.check_status
	FROM
		attendances
	JOIN
		schedules on schedules.id = attendances.schedule_id
	JOIN
		users on attendances.user_id = users.id
	WHERE
		schedules.id = ? AND attendances.check_in is not null AND attendances.check_out is not null
	LIMIT 10 OFFSET ?`, id, offset)
	if err_check != nil {
		return hasil, err_check
	}
	defer result.Close()
	for result.Next() {
		var check entities.GetCheckbyDate
		err := result.Scan(&check.Id, &check.Name, &check.Nik, &check.VaccineStatus, &check.CheckData.Id, &check.CheckData.ScheduleDate, &check.CheckData.Checkin, &check.CheckData.CheckTemperature, &check.CheckData.Checkout, &check.CheckData.CheckStatus)
		if err != nil {
			return hasil, err
		}
		hasil = append(hasil, check)
	}
	return hasil, nil
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
		attendances.id = ? AND date(CONVERT_TZ(schedules.date, '+00:00', '+7:00')) = date(CONVERT_TZ(now(), '+00:00', '+7:00'))`, id)
	if err_check != nil {
		return err_check
	}
	defer result.Close()
	for result.Next() {
		return nil
	}
	return fmt.Errorf("id not found")
}

func (cc *CheckRepository) GetTotalPage() (int, error) {
	var page int
	result := cc.db.QueryRow(`
	SELECT
		count(id)
	FROM
		users 
	WHERE 
		vaccine_status LIKE "%Approved%"`)
	err_scan := result.Scan(&page)
	if err_scan != nil {
		return 0, err_scan
	}
	return int((math.Ceil(float64(page) / float64(10)))), nil
}
