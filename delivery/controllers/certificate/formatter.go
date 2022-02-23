package certificate

type CertificateRequestFormat struct {
	VaccineDose int    `json:"vaccinedose" form:"vaccinedose"`
	Description string `json:"description" form:"description"`
}

type CertificateEditFormat struct {
	Status string `json:"status" form:"status"`
}
