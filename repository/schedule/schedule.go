package schedule

import (
	"database/sql"
	"fmt"
	"math"
	"sirclo/entities"
	"time"
)

type ScheduleRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ScheduleRepository {
	return &ScheduleRepository{db: db}
}

func (sr *ScheduleRepository) CreateSchedule(month time.Month, year int, capacity int, officeId int) error {
	gmt := time.FixedZone("gmt+7", +7*60*60)
	day := 1
	start := time.Date(year, month, day, 0, 0, 0, 0, gmt)
	start_end := time.Date(year, month+1, day, 0, 0, 0, 0, gmt)
	for start != start_end {
		day = day + 1
		start = time.Date(year, month, day, 0, 0, 0, 0, gmt)
		month = start.Month()
		result, err := sr.db.Exec("INSERT INTO schedules (office_id, total_capacity, date) VALUES (?, ?, ?)", officeId, capacity, start)
		if err != nil {
			return err
		}
		mengubah, _ := result.RowsAffected()
		if mengubah == 0 {
			return fmt.Errorf("error gagal terbuat")
		}
	}
	return nil
}

func (sr *ScheduleRepository) EditSchedule(scheduleId int, capacity int) error {
	result, err := sr.db.Exec("UPDATE schedules SET total_capacity = ? WHERE id = ?", capacity, scheduleId)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("schedule not found")
	}
	return nil
}

func (sr *ScheduleRepository) GetSchedule(scheduleId int, offset int) (entities.ScheduleResponse, error) {
	var hasil entities.ScheduleResponse
	var users []entities.UserResponseFormat
	result1, err_users := sr.db.Query(`
	SELECT
		users.id, users.name, users.email, users.image_url, users.nik, offices.Name, users.vaccine_status
	FROM
		users 
	JOIN
		offices ON users.office_id = offices.id
	JOIN
		attendances ON users.id = attendances.user_id
	WHERE 
		attendances.schedule_id = ? AND attendances.status = ?
	ORDER BY
		attendances.updated_at DESC
	LIMIT 10 OFFSET ?`, scheduleId, "approved", offset)
	if err_users != nil {
		return hasil, err_users
	}
	defer result1.Close()
	for result1.Next() {
		var user entities.UserResponseFormat
		err := result1.Scan(&user.Id, &user.Name, &user.Email, &user.ImageUrl, &user.Nik, &user.Office, &user.VaccineStatus)
		if err != nil {
			return hasil, err
		}
		users = append(users, user)
	}
	result2 := sr.db.QueryRow(`
	SELECT
		schedules.id, schedules.date, schedules.total_capacity, schedules.capacity, offices.name
	FROM
		schedules
	JOIN
		offices ON schedules.office_id = offices.id
	WHERE
		schedules.id = ?`, scheduleId)
	err_scan := result2.Scan(&hasil.Id, &hasil.Date, &hasil.TotalCapacity, &hasil.Capacity, &hasil.Office)
	if err_scan != nil {
		return hasil, err_scan
	}

	hasil.Attendance = users
	return hasil, nil
}

func (sr *ScheduleRepository) GetTotalPage(scheduleId int) (int, error) {
	var page int
	result := sr.db.QueryRow(`
	SELECT
		count(users.id)
	FROM
		users 
	JOIN
		attendances ON users.id = attendances.user_id
	WHERE 
		attendances.schedule_id = ?`, scheduleId)
	err_scan := result.Scan(&page)
	if err_scan != nil {
		return 0, err_scan
	}
	return int((math.Ceil(float64(page) / float64(10)))), nil
}
