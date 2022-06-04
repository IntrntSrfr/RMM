package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	"github.com/intrntsrfr/rmm-api/service"
	"golang.org/x/oauth2"
)

type AuthHandler struct {
	r           *gin.Engine
	oauthConfig *oauth2.Config
	j           service.JWTService
}

func NewAuthHandler(r *gin.Engine, o *oauth2.Config, j service.JWTService) {
	h := &AuthHandler{r, o, j}

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

		oc, _ := discordgo.New("Bearer " + token.AccessToken)

		var u *discordgo.User
		res, err := oc.Request("GET", discordgo.EndpointUsers+"@me", nil)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse{CodeError, "Something went wrong"})
			return
		}
		err = json.Unmarshal(res, &u)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse{CodeError, "Something went wrong"})
			return
		}

		jwtToken, err := h.j.GenerateToken(u.ID, token.Expiry, token.AccessToken, token.RefreshToken)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse{CodeError, "Something went wrong"})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": jwtToken,
		})
	}
}
