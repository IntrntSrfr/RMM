package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	r           *gin.Engine
	oauthConfig *oauth2.Config
}

func NewAuthHandler(r *gin.Engine, o *oauth2.Config) {
	h := &AuthHandler{r, o}

	g := h.r.Group("/api/auth")
	g.GET("/callback", h.callback())
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

		token, err := h.oauthConfig.Exchange(context.Background(), code)
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
