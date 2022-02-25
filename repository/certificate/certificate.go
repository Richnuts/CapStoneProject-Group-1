package certificate

import (
	"database/sql"
	"fmt"
	"math"
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

func (cer *CertificateRepository) GetMyCertificate(userId int) (entities.UsersCertificateWithName, error) {
	var hasil entities.UsersCertificateWithName
	var certificates []entities.UsersCertificate
	result, err_certificates := cer.db.Query(`
	SELECT
        certificates.id, certificates.image_url, certificates.vaccine_dose, (select name from users where id = certificates.admin_id) as admin_name, certificates.status, certificates.description
    FROM
        certificates
    JOIN
        users on certificates.user_id = users.id
    WHERE
        certificates.user_id = ?;`, userId)
	if err_certificates != nil {
		return hasil, err_certificates
	}
	defer result.Close()
	for result.Next() {
		var certificate entities.UsersCertificate
		err := result.Scan(&certificate.Id, &certificate.ImageURL, &certificate.VaccineDose, &certificate.AdminName, &certificate.Status, &certificate.Description)
		if err != nil {
			return hasil, err
		}
		hasil.Certificates = certificates
		certificates = append(certificates, certificate)
	}
	result2 := cer.db.QueryRow(`
	SELECT
		users.id, users.name, users.vaccine_status
	FROM
		certificates
	JOIN
		users ON certificates.user_id = users.id
	WHERE
		users.id = ?`, userId)
	err_scan := result2.Scan(&hasil.Id, &hasil.Name, &hasil.Status)
	if err_scan != nil {
		return hasil, err_scan
	}

	hasil.Certificates = certificates
	return hasil, nil
}
func (cer *CertificateRepository) GetUsersCertificates(status string, offset int) ([]entities.UsersCertificateWithName, error) {
	status = "%" + status + "%"
	var hasil []entities.UsersCertificateWithName
	// var certificates []entities.UsersCertificate
	result, err1 := cer.db.Query(`
	SELECT
		id, name, vaccine_status
	FROM
		users
	LIMIT 10 OFFSET ?`, offset)
	if err1 != nil {
		return hasil, err1
	}
	defer result.Close()
	for result.Next() {
		var user entities.UsersCertificateWithName
		err := result.Scan(&user.Id, &user.Name, &user.Status)
		if err != nil {
			return hasil, err
		}
		result1, err_certificates := cer.db.Query(`
	SELECT
        certificates.id, certificates.image_url, certificates.vaccine_dose, (select name from users where id = certificates.admin_id) as admin_name, certificates.status, certificates.description
    FROM
        certificates
    JOIN
        users on certificates.user_id = users.id
    WHERE 
        certificates.status LIKE ? AND users.id = ?
	ORDER BY
		certificates.vaccine_dose ASC`, status, user.Id)
		fmt.Println(status, user.Id)
		if err_certificates != nil {
			return hasil, err_certificates
		}
		defer result1.Close()
		for result1.Next() {
			var certificate entities.UsersCertificate
			err := result1.Scan(&certificate.Id, &certificate.ImageURL, &certificate.VaccineDose, &certificate.AdminName, &certificate.Status, &certificate.Description)
			if err != nil {
				return hasil, err
			}
			user.Certificates = append(user.Certificates, certificate)
		}
		// user.Certificates = certificates
		fmt.Println(user)
		hasil = append(hasil, user)
	}
	return hasil, nil
}

func (cer *CertificateRepository) GetCertificateById(id, userId int) (entities.CertificateResponseGetByIdAndUID, error) {
	var certificate entities.CertificateResponseGetByIdAndUID
	result, err_certificate := cer.db.Query(`
	SELECT
		certificates.id, certificates.image_url, certificates.vaccine_dose, (select name from users where id = certificates.admin_id) as admin_name, certificates.status, certificates.description
	FROM
		certificates
	JOIN
		users as a on certificates.user_id = a.id
	WHERE
		certificates.id = ? AND a.id = ?`, id, userId)
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

func (cer *CertificateRepository) EditCertificate(id, adminId int, status string) error {
	result, err := cer.db.Exec("UPDATE certificates SET status = ?, admin_id = ? WHERE id = ? AND status = ?", status, adminId, id, "Pending")
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("status already assigned")
	}
	return nil
}

func (cer *CertificateRepository) EditMyCertificate(id int, imageURL string) error {
	result, err := cer.db.Exec("UPDATE certificates SET image_url = ?, status = ? WHERE id = ? AND status = ?", imageURL, "Pending", id, "Rejected")
	fmt.Println("anu", imageURL)
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
		user_id = ? AND vaccine_dose = ? AND status != ?`, userId, vaccineDose, "Rejected")
	if err_certificate != nil {
		return err_certificate
	}
	defer result.Close()
	for result.Next() {
		return fmt.Errorf("udah upload atau pending nih")
	}
	return nil
}

func (cer *CertificateRepository) GetVaccineStatus(userId int) error {
	result, err_certificate := cer.db.Query(`
	SELECT
		id
	FROM
		users
	WHERE
		id = ? AND vaccine_status = ?`, userId, "Approved")
	if err_certificate != nil {
		return err_certificate
	}
	defer result.Close()
	for result.Next() {
		return fmt.Errorf("udah sehat, jangan aneh-aneh")
	}
	return nil
}

func (cer *CertificateRepository) GetVaccineDose(id int) (int, error) {
	var dose int
	result := cer.db.QueryRow(`
	SELECT
		vaccine_dose
	FROM
		certificates
	WHERE
		id = ?`, id)
	fmt.Println(id)
	err_scan := result.Scan(&dose)
	if err_scan != nil {
		return dose, err_scan
	}
	return dose, nil
}

func (cer *CertificateRepository) GetTotalPage(status string) (int, error) {
	status = "%" + status + "%"
	var page int
	result := cer.db.QueryRow(`
	SELECT
		count(id)
	FROM
		certificates 
	WHERE 
		status LIKE ?`, status)
	err_scan := result.Scan(&page)
	if err_scan != nil {
		return 0, err_scan
	}
	return int((math.Ceil(float64(page) / float64(10)))), nil
}

func (cer *CertificateRepository) GetName(id int) (string, error) {
	var name string
	result := cer.db.QueryRow(`
	SELECT
		a.name
	FROM
		certificates
	JOIN
		users as a ON certificates.user_id = a.id
	WHERE
		certificates.user_id = ?`, id)
	fmt.Println(id)
	err_scan := result.Scan(&name)
	if err_scan != nil {
		return name, err_scan
	}
	return name, nil
}
