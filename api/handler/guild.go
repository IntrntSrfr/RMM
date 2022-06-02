package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intrntsrfr/rmm-api/service/discord"
)

type GuildHandler struct {
	r *gin.Engine
	d *discord.DiscordService
}

func NewGuildHandler(r *gin.Engine, d *discord.DiscordService) {
	h := &GuildHandler{r, d}
	g := h.r.Group("/api/guilds")
	g.GET("/:guildID", h.getGuild())
	g.GET("/:guildID/members", h.getGuildMembers())
}

func (h *GuildHandler) getGuild() gin.HandlerFunc {
	return func(c *gin.Context) {
		guildID := c.Param("guildID")
		guild, err := h.d.Guild(guildID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{CodeError, err.Error()})
			return
		}
		c.JSON(http.StatusOK, guild)
	}
}

func (h *GuildHandler) getGuildMembers() gin.HandlerFunc {
	return func(c *gin.Context) {
		guildID := c.Param("guildID")
		members, err := h.d.GuildJoins(guildID)
		if err != nil {
			c.JSON(http.StatusBadRequest, ErrorResponse{CodeError, err.Error()})
			return
		}
		c.JSON(http.StatusOK, members)
	}
}
