package certificate

import "sirclo/entities"

type Certificate interface {
	CreateCertificate(userId, vaccineDose int, imageURL, description string) error
	GetMyCertificate(userId int) ([]entities.CertificateResponseGetByIdAndUID, error)
	GetUsersCertificates(status string, offset int) ([]entities.UsersCertificate, error)
	GetCertificateById(id, userId int) (entities.CertificateResponseGetByIdAndUID, error)
	EditCertificate(id int, status string) error
	EditMyCertificate(id int, imageURL string) error
	GetCertificateByDose(userId, vaccineDose int) error
	GetVaccineStatus(userId int) error
	GetVaccineDose(id int) (int, error)
	GetTotalPage(status string) (int, error)
}
