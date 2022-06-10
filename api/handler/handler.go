package handler

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/intrntsrfr/rmm-api/service"

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

type Handler struct {
	e *gin.Engine
}

type Config struct {
	Discord     *discord.DiscordService
	OauthConfig *oauth2.Config
	JwtUtil     *service.JWTUtil
}

func NewHandler(conf *Config) *Handler {
	h := &Handler{gin.Default()}
	h.e.Use(Cors())

	NewAuthHandler(h.e, conf.OauthConfig, conf.JwtUtil)
	NewGuildHandler(h.e, conf.Discord, conf.JwtUtil)

	h.e.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return h
}

func (h *Handler) Run(address string) error {
	return h.e.Run(address)
}

func Cors() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	})
}
