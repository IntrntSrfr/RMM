package service

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
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

type Claims struct {
	jwt.RegisteredClaims
	Token        string `json:"token"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

func (j *JWTUtil) GenerateToken(id string, expiry time.Time, token string, refreshToken string) (string, error) {
	tkn := jwt.New(jwt.SigningMethodHS256)
	claims := Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "RMM",
			Subject:   id,
			ExpiresAt: jwt.NewNumericDate(expiry),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		Token:        token,
		TokenType:    "Bearer",
		RefreshToken: refreshToken,
	}

	tkn.Claims = claims
	return tkn.SignedString(j.key)
}

func (j *JWTUtil) IsAuthorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(strings.ToLower(auth), "bearer") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}
		tokenString := strings.TrimPrefix(auth, "Bearer ")
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
			return j.key, nil
		})
		if err != nil {
			fmt.Println(err)
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "mangled token"})
			return
		}
		if !token.Valid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
			return
		}

		c.Set("claims", token.Claims)
		c.Next()
	}
}
