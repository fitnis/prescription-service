package main

import (
	"github.com/fitnis/prescription-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	p := router.Group("/prescriptions")
	{
		p.POST("", handlers.CreatePrescription)
		p.POST("/prescribe", handlers.PrescribeMedication)
	}

	router.Run(":8086")
}
