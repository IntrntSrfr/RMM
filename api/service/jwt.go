package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"time"
)

type JWTService interface {
	ValidateToken(token string) (*jwt.Token, error)
	GenerateToken(id string, expiry time.Time, token string, refreshToken string) (string, error)
	JWT() gin.HandlerFunc
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

func (j *JWTUtil) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return j.key, nil
	})
}

func (j *JWTUtil) JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		if token, err := j.ValidateToken(token); err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		} else {
			c.Set("token", token)
			claims, _ := token.Claims.(jwt.MapClaims)
			c.Set("claims", claims)
		}
		c.Next()
	}
}
