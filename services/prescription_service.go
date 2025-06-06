package services

import (
	"github.com/fitnis/prescription-service/models"
)

var prescriptions []models.Prescription

func PrescribeMedication(p models.Prescription) models.GenericResponse {
	prescriptions = append(prescriptions, p)
	return models.GenericResponse{Message: "Medication prescribed"}
}
