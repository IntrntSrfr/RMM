package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Code int

const (
	CodeOK Code = 1 << iota
	CodeError
)

type ErrorResponse struct {
	Code    Code   `json:"code"`
	Message string `json:"message"`
}

func NewHandler() *gin.Engine {
	r := gin.Default()
	r.Use(Cors())

	NewAuthHandler(r)
	r.GET("/api/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	return r
}
