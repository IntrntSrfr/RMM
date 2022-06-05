package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
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

func (j *JWTUtil) ValidateToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})
	if err != nil {
		return nil, nil, err
	}
	return token, claims, nil
}

func (j *JWTUtil) JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			fmt.Println("no token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		// remove "Bearer "
		tokenString = tokenString[7:]

		token, claims, err := j.ValidateToken(tokenString)
		if err != nil {
			fmt.Println(err)
			fmt.Println("token not valid")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
			return
		}

		fmt.Println("setting token")
		c.Set("token", token)
		c.Set("claims", claims)
		c.Next()
	}
}
