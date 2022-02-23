package certificate

import "sirclo/entities"

type Certificate interface {
	CreateCertificate(userId, vaccineDose int, imageURL, description string) error
	GetApprovedCertificate(userId int) ([]entities.CertificateResponseGetByIdAndUID, error)
	GetCertificateById(id, userId int) (entities.CertificateResponseGetByIdAndUID, error)
	EditCertificate(id int, status string) error
	GetCertificateByDose(userId, vaccineDose int) error
}
