package certificate

import (
	"database/sql"
	"fmt"
)

type CertificateRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CertificateRepository {
	return &CertificateRepository{db: db}
}

func (cer *CertificateRepository) CreateCertificate(userId int, imageURL, description string) error {
	result, err := cer.db.Exec("INSERT INTO certificates (user_id, image_url, description) VALUES (?, ?, ?)", userId, imageURL, description)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("error gagal terbuat")
	}
	return nil
}
