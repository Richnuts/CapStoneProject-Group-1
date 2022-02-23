package certificate

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type CertificateRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CertificateRepository {
	return &CertificateRepository{db: db}
}

func (cer *CertificateRepository) CreateCertificate(userId, vaccineDose int, imageURL, description string) error {
	result, err := cer.db.Exec("INSERT INTO certificates (user_id, image_url, status, vaccine_dose, description) VALUES (?, ?, ?, ?, ?)", userId, imageURL, "Pending", vaccineDose, description)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("failed to create data")
	}
	return nil
}

func (cer *CertificateRepository) GetApprovedCertificate(userId int) ([]entities.CertificateResponseGetByIdAndUID, error) {
	var certificates []entities.CertificateResponseGetByIdAndUID
	result, err_certificates := cer.db.Query(`
	SELECT
		certificates.id, certificates.image_url, certificates.vaccine_dose, certificates.status, certificates.description
	FROM
		certificates
	JOIN
		users on certificates.user_id = users.id
	WHERE
		users.id = ? AND certificates.status != ?`, userId, "Rejected")
	if err_certificates != nil {
		return certificates, err_certificates
	}
	defer result.Close()
	for result.Next() {
		var certificate entities.CertificateResponseGetByIdAndUID
		err := result.Scan(&certificate.Id, &certificate.ImageURL, &certificate.VaccineDose, &certificate.Status, &certificate.Description)
		if err != nil {
			return certificates, err
		}
		certificates = append(certificates, certificate)
	}
	return certificates, nil
}

func (cer *CertificateRepository) GetCertificateById(id, userId int) (entities.CertificateResponseGetByIdAndUID, error) {
	var certificate entities.CertificateResponseGetByIdAndUID
	result, err_certificate := cer.db.Query(`
	SELECT
		certificates.id, certificates.image_url, certificates.vaccine_dose, certificates.status, certificates.description
	FROM
		certificates
	JOIN
		users on certificates.user_id = users.id
	WHERE
		certificates.id = ? AND users.id = ?`, id, userId)
	if err_certificate != nil {
		return certificate, err_certificate
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&certificate.Id, &certificate.ImageURL, &certificate.VaccineDose, &certificate.Status, &certificate.Description)
		if err != nil {
			return certificate, err
		}
	}
	return certificate, nil
}

func (cer *CertificateRepository) EditCertificate(id int, status string) error {
	result, err := cer.db.Exec("UPDATE certificates SET status = ? WHERE id = ? AND status = ?", status, id, "Pending")
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("status already assigned")
	}
	return nil
}

func (cer *CertificateRepository) GetCertificateByDose(userId, vaccineDose int) error {
	result, err_certificate := cer.db.Query(`
	SELECT
		id
	FROM
		certificates
	WHERE
		user_id = ? AND vaccine_dose = ?`, userId, vaccineDose)
	if err_certificate != nil {
		return err_certificate
	}
	defer result.Close()
	for result.Next() {
		return fmt.Errorf("udah upload atau pending nih")
	}
	return nil
}
