package handler

import (
	"net/http"

	"github.com/intrntsrfr/rmm-api/service/discord"
	"golang.org/x/oauth2"

	"github.com/gin-gonic/gin"
)

type Code int

const (
	CodeError Code = 1 << iota
)

type ErrorResponse struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

type Config struct {
	Discord     *discord.DiscordService
	OauthConfig *oauth2.Config
}

func NewHandler(conf *Config) *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	NewAuthHandler(r, conf.OauthConfig)
	NewGuildHandler(r, conf.Discord)

	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
