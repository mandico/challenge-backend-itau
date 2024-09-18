package main

import (
	"github.com/gin-gonic/gin"
	"github.com/mandico/challenge-backend/internal/controller"
	"github.com/mandico/challenge-backend/internal/service"
)

func main() {
	r := gin.Default()

	jwtService := service.NewJwtService("itau")
	jwtController := controller.NewJwtController(jwtService)

	r.GET("/validate", jwtController.ValidateJwt)

	r.Run(":8888")
}
