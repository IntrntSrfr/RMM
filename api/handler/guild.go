package handler

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/intrntsrfr/rmm-api/service"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/intrntsrfr/rmm-api/service/discord"
)

type GuildHandler struct {
	r       *gin.Engine
	d       *discord.DiscordService
	jwtUtil *service.JWTUtil
}

func NewGuildHandler(r *gin.Engine, d *discord.DiscordService, j *service.JWTUtil) {
	h := &GuildHandler{r, d, j}
	g := h.r.Group("/api/guilds")
	g.GET("/", j.JWT(), h.getGuilds())
	g.GET("/:guildID", h.getGuild())
	g.GET("/:guildID/members", h.getGuildMembers())
}

func (h *GuildHandler) getGuilds() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("guild start")
		claims := c.MustGet("claims").(*service.Claims)
		fmt.Println("claims")
		oc, _ := discordgo.New("Bearer " + claims.Token)
		fmt.Println("new token state")
		res, err := oc.Request("GET", discordgo.EndpointUsers+"@me/guilds", nil)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse{CodeError, "Something went wrong"})
			return
		}
		var gs []*discordgo.Guild
		err = json.Unmarshal(res, &gs)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusInternalServerError, ErrorResponse{CodeError, "Something went wrong"})
			return
		}

		var mutualGs []*discordgo.Guild
		for _, g := range gs {
			_, err := h.d.Guild(g.ID)
			if err != nil {
				continue
			}
			if g.Permissions&discordgo.PermissionBanMembers == 0 {
				continue
			}
			mutualGs = append(mutualGs, g)
		}
		c.JSON(http.StatusOK, mutualGs)
	}
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
