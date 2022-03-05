package certificate

type CertificateRequestFormat struct {
	VaccineDose int    `json:"vaccine_dose" form:"vaccine_dose"`
	Description string `json:"description" form:"description"`
}

type CertificateEditFormat struct {
	Status string `json:"status" form:"status"`
}
