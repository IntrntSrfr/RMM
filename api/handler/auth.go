package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	r *gin.Engine
}

func NewAuthHandler(r *gin.Engine) {
	ah := &AuthHandler{r}

	g := ah.r.Group("/api/auth")
	g.GET("/callback", ah.callback())
}

func (h *AuthHandler) callback() gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		if code == "" {
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:    CodeError,
				Message: "no code provided",
			})
			return
		}

		conf := &oauth2.Config{
			ClientID:     "394162399348785152",
			ClientSecret: "OamPZi012BJ8BBSP7AM1PgDwfb8s9CTB",
			Scopes:       []string{},
			Endpoint: oauth2.Endpoint{
				TokenURL: "https://discord.com/api/oauth2/token",
			},
			RedirectURL: "http://localhost:3000/callback",
		}

		token, err := conf.Exchange(context.Background(), code)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, ErrorResponse{
				Code:    CodeError,
				Message: "could not exchange token",
			})
			return
		}

		c.JSON(http.StatusOK, token)
	}
}
