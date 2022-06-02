package service

import (
	"github.com/golang-jwt/jwt/v4"
)

type JWTService interface {
	GenerateToken() (string, error)
	ValidateToken(token string)
}

type JWTUtil struct {
	key []byte
}

func NewJWTUtil(key []byte) *JWTUtil {
	return &JWTUtil{key}
}

func (j *JWTUtil) GenerateToken() (string, error) {
	tkn := jwt.New(jwt.SigningMethodHS256)
	tkn.Claims = jwt.RegisteredClaims{
		Issuer: "RMM",
		Subject: "Discord User",
		
	}
	return jwt.New(jwt.SigningMethodHS256).SignedString(j.key)
}

func (j *JWTUtil) ValidateToken(token string) {
	// TODO implement me
	panic("implement me")
}
