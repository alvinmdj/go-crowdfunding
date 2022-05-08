package auth

import (
	"errors"
	"go-crowdfunding/helper"
	"os"

	"github.com/golang-jwt/jwt"
)

type Service interface {
	GenerateToken(userId int) (string, error)
	ValidateToken(encodedToken string) (*jwt.Token, error)
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

func (s *jwtService) ValidateToken(encodedToken string) (*jwt.Token, error) {
	token, err := jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, errors.New("invalid token")
		}

		helper.LoadEnv()
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return token, err
	}

	return token, nil
}
