package handlers

import (
	"net/http"

	"github.com/fitnis/prescription-service/models"
	"github.com/fitnis/prescription-service/services"
	"github.com/gin-gonic/gin"
)

func RegisterPrescriptionRoutes(rg *gin.RouterGroup) {
	p := rg.Group("/prescriptions")
	p.POST("", prescribe)
	p.POST("/prescribe", prescribe)
}

func prescribe(c *gin.Context) {
	var req models.Prescription
	_ = c.ShouldBindJSON(&req)
	c.JSON(http.StatusCreated, services.PrescribeMedication(req))
}
