package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mandico/challenge-backend/internal/controller"
	"github.com/mandico/challenge-backend/internal/service"
)

func main() {
	r := gin.Default()

	jwtService := service.NewJwtService("itau")
	jwtController := controller.NewJwtController(jwtService)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	})

	r.GET("/validate", jwtController.ValidateJwt)

	r.Run(":8888")
}
