package certificate

import "sirclo/entities"

type Certificate interface {
	CreateCertificate(userId, vaccineDose int, imageURL, description string) error
	GetMyCertificate(userId int) (entities.UsersCertificateWithName, error)
	GetUsersCertificates(status string, offset int) ([]entities.UsersCertificateWithName, error)
	GetCertificateById(id, userId int) (entities.CertificateResponseGetByIdAndUID, error)
	EditCertificate(id, adminId int, status string) error
	EditMyCertificate(id int, imageURL string) error
	GetCertificateByDose(userId, vaccineDose int) error
	GetVaccineStatus(userId int) error
	GetVaccineDose(id int) (int, error)
	GetTotalPage(status string) (int, error)
	GetTotalUsers() (int, error)
	GetName(id int) (string, error)
}
