package attendance

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
	"time"
)

type AttendanceRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

func (ar AttendanceRepository) CreateAttendance(userId int, scheduleId int, description string, imageUrl string) error {
	result, err := ar.db.Exec("INSERT INTO attendances (user_id, schedule_id, description, image_url) VALUES (?, ?, ?, ?)", userId, scheduleId, description, imageUrl)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("error gagal terbuat")
	}
	return nil
}

func (ar AttendanceRepository) GetUserAttendanceStatus(userId int, scheduleId int) error {
	result, err := ar.db.Query("SELECT id FROM attendances WHERE status != ? AND user_id = ? AND schedule_id = ?", "Rejected", userId, scheduleId)
	if err != nil {
		return err
	}
	defer result.Close()
	for result.Next() {
		return fmt.Errorf("request sudah ada")
	}
	return nil
}

func (ar AttendanceRepository) GetUserVaccineStatus(userId int) error {
	result, err := ar.db.Query(`
	SELECT 
		id
	FROM 
		users
	WHERE 
		id = ? AND vaccine_status = ?`, userId, "Approved")
	defer result.Close()
	if err != nil {
		return err
	}
	for result.Next() {
		return nil
	}
	return fmt.Errorf("user belom vaksin")
}

func (ar AttendanceRepository) EditAttendance(attendanceId int, adminId int, status string, statusInfo string) error {
	result, err := ar.db.Exec("UPDATE attendances SET admin_id = ?, status = ?, status_info = ?, updated_at = now() WHERE id = ? AND status = ?", adminId, status, statusInfo, attendanceId, "Pending")
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("error gagal mengapprove")
	}
	return nil
}

func (ar AttendanceRepository) GetPendingAttendance(offset int, officeId int) ([]entities.PendingAttendance, error) {
	var hasilAkhir []entities.PendingAttendance
	result, err_query := ar.db.Query(`
	SELECT 
		attendances.id, CONVERT_TZ(schedules.date, '+00:00', '+7:00'), attendances.schedule_id, (select (total_capacity-capacity) from schedules where id = attendances.schedule_id), offices.name, attendances.image_url, attendances.description, CONVERT_TZ(attendances.created_at, '+00:00', '+7:00'), attendances.user_id, users.name, users.email, users.image_url, users.nik, users.vaccine_status, (select name from offices where id = users.office_id)
	FROM 
		attendances 
	JOIN
		schedules ON attendances.schedule_id = schedules.id
	JOIN
		offices ON schedules.office_id = offices.id 
	JOIN
		users ON attendances.user_id = users.id
	WHERE 
		attendances.status = ? AND schedules.office_id = ?
	ORDER BY attendances.created_at ASC
	LIMIT 10 OFFSET ?`, "Pending", officeId, offset)
	defer result.Close()
	if err_query != nil {
		return hasilAkhir, fmt.Errorf("request wfo not found")
	}
	for result.Next() {
		var hasil entities.PendingAttendance
		err := result.Scan(&hasil.Id, &hasil.Date, &hasil.ScheduleId, &hasil.ActualCapacity, &hasil.Office, &hasil.ImageUrl, &hasil.Description, &hasil.RequestTime, &hasil.User.Id, &hasil.User.Name, &hasil.User.Email, &hasil.User.ImageUrl, &hasil.User.Nik, &hasil.User.VaccineStatus, &hasil.User.Office)
		if err != nil {
			return hasilAkhir, err
		}
		hasilAkhir = append(hasilAkhir, hasil)
	}
	return hasilAkhir, nil
}

func (ar AttendanceRepository) GetPendingAttendanceTotalData(officeId int) (int, error) {
	var hasil int
	result := ar.db.QueryRow(`
	SELECT 
		count(attendances.id)
	FROM
		attendances
	JOIN
		schedules ON attendances.schedule_id = schedules.id
	WHERE
		attendances.status = ? AND schedules.office_id = ?`, "Pending", officeId)
	err := result.Scan(&hasil)
	if err != nil {
		return hasil, err
	}
	return hasil, nil
}

func (ar AttendanceRepository) CheckCapacity(scheduleId int) (int, error) {
	var hasil int
	result := ar.db.QueryRow("select (total_capacity - capacity) from schedules where id = ?", scheduleId)
	err := result.Scan(&hasil)
	if err != nil {
		return hasil, err
	}
	return hasil, nil
}

func (ar AttendanceRepository) GetAttendanceById(attendanceId int) (entities.AttendanceGetFormat, error) {
	var hasil entities.AttendanceGetFormat
	result := ar.db.QueryRow(`
	SELECT 
		attendances.id, users.name, CONVERT_TZ(schedules.date, '+00:00', '+7:00'), offices.name, attendances.status,  attendances.status_info, (select name from users where id = attendances.admin_id) AS checker, CONVERT_TZ(attendances.check_in, '+00:00', '+7:00')
	FROM 
		attendances 
	JOIN
		schedules ON attendances.schedule_id = schedules.id
	JOIN
		offices ON schedules.office_id = offices.id 
	JOIN
		users ON attendances.user_id = users.id
	WHERE 
		attendances.id = ?`, attendanceId)
	err := result.Scan(&hasil.Id, &hasil.Name, &hasil.Date, &hasil.Office, &hasil.Status, &hasil.StatusInfo, &hasil.AdminName, &hasil.CheckIn)
	if err != nil {
		return hasil, err
	}
	return hasil, nil
}

func (ar AttendanceRepository) GetMyAttendance(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error) {
	var hasilAkhir []entities.AttendanceGetFormat
	status = "%" + status + "%"
	result, err_query := ar.db.Query(`
	SELECT 
		attendances.id, users.name, CONVERT_TZ(schedules.date, '+00:00', '+7:00'), offices.name, attendances.status, attendances.status_info, (select name from users where id = attendances.admin_id) AS checker, CONVERT_TZ(attendances.check_in, '+00:00', '+7:00'), CONVERT_TZ(attendances.created_at, '+00:00', '+7:00'), CONVERT_TZ(attendances.updated_at, '+00:00', '+7:00')
	FROM 
		attendances 
	JOIN
		schedules ON attendances.schedule_id = schedules.id
	JOIN
		offices ON schedules.office_id = offices.id 
	JOIN
		users ON attendances.user_id = users.id
	WHERE 
		attendances.user_id = ? AND attendances.status like ? LIMIT 10 OFFSET ?`, userId, status, offset)
	defer result.Close()
	if err_query != nil {
		return hasilAkhir, fmt.Errorf("request wfo not found")
	}

	for result.Next() {
		var hasil entities.AttendanceGetFormat
		err := result.Scan(&hasil.Id, &hasil.Name, &hasil.Date, &hasil.Office, &hasil.Status, &hasil.StatusInfo, &hasil.AdminName, &hasil.CheckIn, &hasil.RequestTime, &hasil.ApprovedTime)
		if err != nil {
			fmt.Println(err)
			return hasilAkhir, err
		}
		hasilAkhir = append(hasilAkhir, hasil)
	}
	return hasilAkhir, nil
}

func (ar AttendanceRepository) GetMyAttendanceTotalData(userId int, status string) (int, error) {
	var hasil int
	status = "%" + status + "%"
	result := ar.db.QueryRow(`
	SELECT 
		count(id)
	FROM 
		attendances 
	WHERE 
		user_id = ? AND status LIKE ?`, userId, status)
	err := result.Scan(&hasil)
	if err != nil {
		return hasil, err
	}
	return hasil, nil
}

func (ar AttendanceRepository) GetMyAttendanceSortByLatest(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error) {
	var hasilAkhir []entities.AttendanceGetFormat
	status = "%" + status + "%"
	result, err_query := ar.db.Query(`
	SELECT 
		attendances.id, users.name, CONVERT_TZ(schedules.date, '+00:00', '+7:00'), offices.name, attendances.status, attendances.status_info, (select name from users where id = attendances.admin_id) AS checker, CONVERT_TZ(attendances.check_in, '+00:00', '+7:00'), CONVERT_TZ(attendances.created_at, '+00:00', '+7:00'), CONVERT_TZ(attendances.updated_at, '+00:00', '+7:00')
	FROM 
		attendances 
	JOIN
		schedules ON attendances.schedule_id = schedules.id
	JOIN
		offices ON schedules.office_id = offices.id 
	JOIN
		users ON attendances.user_id = users.id
	WHERE 
		attendances.user_id = ? AND attendances.status like ?
	ORDER BY attendances.created_at ASC LIMIT 10 OFFSET ?`, userId, status, offset)
	defer result.Close()
	if err_query != nil {
		return hasilAkhir, fmt.Errorf("request wfo not found")
	}

	for result.Next() {
		var hasil entities.AttendanceGetFormat
		err := result.Scan(&hasil.Id, &hasil.Name, &hasil.Date, &hasil.Office, &hasil.Status, &hasil.StatusInfo, &hasil.AdminName, &hasil.CheckIn, &hasil.RequestTime, &hasil.ApprovedTime)
		if err != nil {
			return hasilAkhir, err
		}
		hasilAkhir = append(hasilAkhir, hasil)
	}
	return hasilAkhir, nil
}

func (ar AttendanceRepository) GetMyAttendanceSortByLongest(userId int, offset int, status string) ([]entities.AttendanceGetFormat, error) {
	var hasilAkhir []entities.AttendanceGetFormat
	status = "%" + status + "%"
	result, err_query := ar.db.Query(`
	SELECT 
		attendances.id, users.name, CONVERT_TZ(schedules.date, '+00:00', '+7:00'), offices.name, attendances.status, attendances.status_info, (select name from users where id = attendances.admin_id) AS checker, CONVERT_TZ(attendances.check_in, '+00:00', '+7:00'), CONVERT_TZ(attendances.created_at, '+00:00', '+7:00'), CONVERT_TZ(attendances.updated_at, '+00:00', '+7:00')
	FROM 
		attendances 
	JOIN
		schedules ON attendances.schedule_id = schedules.id
	JOIN
		offices ON schedules.office_id = offices.id 
	JOIN
		users ON attendances.user_id = users.id
	WHERE 
		attendances.user_id = ? AND attendances.status like ?
	ORDER BY schedules.date DESC LIMIT 10 OFFSET ?`, userId, status, offset)
	defer result.Close()
	if err_query != nil {
		return hasilAkhir, fmt.Errorf("request wfo not found")
	}

	for result.Next() {
		var hasil entities.AttendanceGetFormat
		err := result.Scan(&hasil.Id, &hasil.Name, &hasil.Date, &hasil.Office, &hasil.Status, &hasil.StatusInfo, &hasil.AdminName, &hasil.CheckIn, &hasil.RequestTime, &hasil.ApprovedTime)
		if err != nil {
			return hasilAkhir, err
		}
		hasilAkhir = append(hasilAkhir, hasil)
	}
	return hasilAkhir, nil
}

func (ar AttendanceRepository) CheckCreateRequestDate(scheduleId int) (time.Time, error) {
	result := ar.db.QueryRow(`SELECT
        schedules.date
    FROM
        schedules
    WHERE
        schedules.id = ?`, scheduleId)
	var hasil time.Time
	err := result.Scan(&hasil)
	if err != nil {
		return hasil, err
	}
	return hasil, nil
}
