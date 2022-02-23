package attendance

import (
	"database/sql"
	"fmt"
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
	result, err := ar.db.Query("SELECT id FROM attendances WHERE status != ? AND user_id = ? AND schedule_id = ?", "rejected", userId, scheduleId)
	if err != nil {
		return err
	}
	defer result.Close()
	for result.Next() {
		return fmt.Errorf("request sudah ada")
	}
	return nil
}
