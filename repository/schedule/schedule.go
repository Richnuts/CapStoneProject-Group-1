package schedule

import (
	"database/sql"
	"fmt"
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

func (sr *ScheduleRepository) EditSchedule(date string, capacity int, officeId int) error {
	result, err := sr.db.Exec("UPDATE schedules SET total_capacity = ? WHERE date = ? AND office_id = ?", capacity, date, officeId)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("schedule not found")
	}
	return nil
}
