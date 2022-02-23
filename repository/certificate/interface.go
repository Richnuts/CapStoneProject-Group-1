package certificate

import "sirclo/entities"

type Certificate interface {
	CreateCertificate(userId int, imageURL, description string) error
	GetCertificate(userId int) (entities.CertificateResponse, error)
	GetCertificateById(id, userId int) (entities.CertificateResponse, error)
	EditCertificate(certificate entities.Certificate) error
}
