package service

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTService interface {
	ValidateToken(token string) bool
	GenerateToken(id string, expiry time.Time, token string, refreshToken string) (string, error)
}

type JWTUtil struct {
	key []byte
}

func NewJWTUtil(key []byte) *JWTUtil {
	return &JWTUtil{key}
}

func (j *JWTUtil) GenerateToken(id string, expiry time.Time, token string, refreshToken string) (string, error) {
	tkn := jwt.New(jwt.SigningMethodHS256)
	claims := jwt.MapClaims{
		"iss": "rmm",
		"sub": id,
		"exp": jwt.NewNumericDate(expiry),
		"iat": jwt.NewNumericDate(time.Now()),
		"tkn": token,
		"ttp": "Bearer",
		"rsh": refreshToken,
	}
	tkn.Claims = claims
	return tkn.SignedString(j.key)
}

func (j *JWTUtil) ValidateToken(token string) bool {
	// TODO implement me
	panic("implement me")
}
