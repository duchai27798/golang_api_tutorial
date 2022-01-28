package service

import (
	"fmt"
	"github.com/duchai27798/golang_api_tutorial/src/utils"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type IJWTService interface {
	GenerateToken(UserId string) string
	ValidateToken(token string) (*jwt.Token, error)
}

type JWTCustomClaim struct {
	UserID string `json:"user_id"`
	jwt.StandardClaims
}

type JWTService struct {
	secretKey string
	issuer    string
}

func (j JWTService) GenerateToken(UserId string) string {
	claims := &JWTCustomClaim{
		UserId,
		jwt.StandardClaims{
			ExpiresAt: time.Now().AddDate(0, 0, 1).Unix(),
			Issuer:    j.issuer,
			IssuedAt:  time.Now().Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		utils.LogObj(err)
	}
	return t
}

func (j JWTService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method %v", t.Header["qlg"])
		}
		return []byte(j.secretKey), nil
	})
}

func NewJWTService() IJWTService {
	return &JWTService{
		issuer:    "harry",
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	return os.Getenv("JWT_SECRET")
}
