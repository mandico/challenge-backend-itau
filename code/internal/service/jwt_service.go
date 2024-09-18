package service

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"

	"github.com/dgrijalva/jwt-go"
)

var mySigningKey = []byte("itau")

type JwtService struct {
	secretKey string
}

func NewJwtService(secretKey string) *JwtService {
	return &JwtService{secretKey: secretKey}
}

func (s *JwtService) ValidateJwt(tokenString string) (bool, error) {

	token, err := parseJWT(tokenString)
	if err != nil {
		fmt.Println("Erro ao analisar o token:", err)
		return false, err
	}

	claims, err := extractClaims(token)
	if err != nil {
		fmt.Println("Erro ao extrair as claims:", err)
		return false, err
	}

	// validate claim 'Name'
	name, ok := claims["Name"].(string)
	if !ok || !regexp.MustCompile(`^[a-zA-Z ]{1,256}$`).MatchString(name) || len(name) > 256 {
		return false, fmt.Errorf("claim 'Nome' inválido")
	}

	// validate claim 'Role'
	validRoles := map[string]bool{"Admin": true, "Member": true, "External": true}
	role, ok := claims["Role"].(string)
	if !ok || !validRoles[role] {
		return false, fmt.Errorf("claim 'Role' inválido")
	}

	// validate exist claims diferent of 'Name' and 'Role' and 'Seed'
	if !hasMoreThanThreeClaims(claims) {
		return false, fmt.Errorf("há mais de 3 claims")
	}

	// validate claim 'Seed'
	seedStr, ok := claims["Seed"].(string)
	if !ok {
		return false, fmt.Errorf("claim 'Seed' não é uma string")
	}
	seed, err := strconv.Atoi(seedStr)
	if err != nil {
		return false, fmt.Errorf("erro ao converter claim 'Seed' para int: %v", err)
	}
	if !ok || !isPrime(seed) {
		return false, fmt.Errorf("seed não é um número primo")
	}
	return true, nil
}

func parseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("método de assinatura inesperado: %v", token.Method)
		}
		return mySigningKey, nil

	})

	if err != nil {
		return nil, err
	}

	return token, nil
}

// extractClaims :: valida o token e retorna as claims
func extractClaims(token *jwt.Token) (map[string]interface{}, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, errors.New("token inválido")
	}
	return claims, nil
}

// hasMoreThanThreeClaims :: verifica se o token possui mais de 3 claims
func hasMoreThanThreeClaims(claims map[string]interface{}) bool {
	expectedClaims := []string{"Name", "Role", "Seed"}
	status := true
	for claim := range claims {
		found := false
		for _, expected := range expectedClaims {
			if claim == expected {
				found = true
				break
			}
		}
		if !found {
			status = false
		}
	}
	return status

}
