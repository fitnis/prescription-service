package models

type Prescription struct {
	PatientID  string `json:"patientId"`
	Medication string `json:"medication"`
	Dosage     string `json:"dosage"`
}

type GenericResponse struct {
	Message string `json:"message"`
}
