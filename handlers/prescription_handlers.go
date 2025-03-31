package handlers

import (
	"net/http"

	"github.com/fitnis/prescription-service/models"
	"github.com/fitnis/prescription-service/services"
	"github.com/gin-gonic/gin"
)

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","medication":"Ibuprofen","dosage":"200mg"}' http://localhost:8086/prescriptions
func CreatePrescription(c *gin.Context) {
	var req models.Prescription
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	c.JSON(http.StatusCreated, services.CreatePrescription(req))
}

// curl -X POST -H "Content-Type: application/json" -d '{"patientId":"123","medication":"Ibuprofen","dosage":"200mg"}' http://localhost:8086/prescriptions/prescribe
func PrescribeMedication(c *gin.Context) {
	var req models.Prescription
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}
	c.JSON(http.StatusCreated, services.PrescribeMedication(req))
}
