package main

import (
	"github.com/fitnis/prescription-service/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	p := router.Group("/prescriptions")
	{
		handlers.RegisterPrescriptionRoutes(p)
	}

	router.Run(":8086")
}
