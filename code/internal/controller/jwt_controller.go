package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mandico/challenge-backend/internal/service"
)

type JwtController struct {
	jwtService *service.JwtService
}

func NewJwtController(jwtService *service.JwtService) *JwtController {
	return &JwtController{jwtService: jwtService}
}

func (c *JwtController) ValidateJwt(ctx *gin.Context) {
	token := ctx.Query("jwt")
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "token não informado", "result": false})
		return
	}

	result, err := c.jwtService.ValidateJwt(token)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": err.Error(), "result": result})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "token válido", "result": result})
}
