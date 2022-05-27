package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/intrntsrfr/rmm-api/service/discord"
	"net/http"
)

type GuildHandler struct {
	r *gin.Engine
	d *discord.DiscordService
}

func NewGuildHandler(r *gin.Engine, d *discord.DiscordService) {
	h := &GuildHandler{r, d}
	g := h.r.Group("/api/guilds")
	g.GET("/:guildID", h.getGuild())
}

func (h *GuildHandler) getGuild() gin.HandlerFunc {
	return func(c *gin.Context) {
		guildID := c.Param("guildID")
		members, err := h.d.RecentMembers(guildID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, ErrorResponse{
				Code:    CodeError,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, members)
	}
}
