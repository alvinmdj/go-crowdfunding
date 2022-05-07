package auth

import (
	"go-crowdfunding/helper"
	"os"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct{}

func NewService() *jwtService {
	return &jwtService{}
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	helper.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
