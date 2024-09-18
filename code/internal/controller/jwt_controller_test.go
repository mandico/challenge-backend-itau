package controller

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/mandico/challenge-backend/internal/service"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	jwtService := service.NewJwtService("itau")
	jwtController := NewJwtController(jwtService)
	r.GET("/validate", jwtController.ValidateJwt)
	return r
}

func TestValidateJwt(t *testing.T) {
	router := setupRouter()

	tests := []struct {
		name       string
		jwt        string
		statusCode int
	}{
		{
			name:       "Valid JWT",
			jwt:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJSb2xlIjoiQWRtaW4iLCJTZWVkIjoiNzg0MSIsIk5hbWUiOiJUb25pbmhvQXJhdWpvIn0.tGpW30PdezwBOSoLEMy54PwUen4TwZ243ke_EOC75CA",
			statusCode: http.StatusOK,
		},
		{
			name:       "Invalid JWT",
			jwt:        "eyJhbGciOiJzI1NiJ9.dfsdfsfryJSr2xrIjoiQWRtaW4iLCJTZrkIjoiNzg0MSIsIk5hbrUiOiJUb25pbmhvIEFyYXVqbyJ9.QY05fsdfsIjtrcJnP533kQNk8QXcaleJ1Q01jWY_ZzIZuAg",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "JWT with invalid Claim Name",
			jwt:        "eyJhbGciOiJIUzI1NiJ9.eyJSb2xlIjoiRXh0ZXJuYWwiLCJTZWVkIjo4ODAzNywiTmFtZSI6Ik00cmlhIE9saXZpYSJ9.M24ssnXg3xtb5F0rvWUQyckc7JAA-RUY-8XXBo41138",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "JWT with more than 3 claims",
			jwt:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJSb2xlIjoiQWRtaW4iLCJTZWVkIjo3ODQxLCJPcmciOiJCUiIsIk5hbWUiOiJUb25pbmhvIEFyYXVqbyJ9.qADMPYJaRFwAatSsHGExu5H3BhaPYRsKmMLLa9FUcKg",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "JWT with invalid Role",
			jwt:        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJSb2xlIjoiQmFuYW5hIiwiU2VlZCI6Nzg0MSwiTmFtZSI6IlRvbmluaG8gQXJhdWpvIn0.ZRgnjhurQDfo63j_FOjEGOw_bV72OIjaHZuz6Uz1ERM",
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req, _ := http.NewRequest("GET", "/validate?jwt="+tt.jwt, nil)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.statusCode, w.Code)
		})
	}
}
