package auth

import (
	"fmt"
	"go-crowdfunding/helper"
	"os"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId int) (string, error)
}

type jwtService struct {
}

func (s *jwtService) GenerateToken(userId int) (string, error) {
	claim := jwt.MapClaims{}
	claim["user_id"] = userId

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	helper.LoadEnv()
	jwtSecret := os.Getenv("JWT_SECRET")
	fmt.Println("call jwt secret env from jwt service:", jwtSecret)

	signedToken, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return signedToken, err
	}

	return signedToken, nil
}
