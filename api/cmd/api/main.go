package main

import (
	"encoding/json"
	"github.com/intrntsrfr/rmm-api/service"
	"io/ioutil"

	"github.com/intrntsrfr/rmm-api/service/discord"

	"github.com/intrntsrfr/rmm-api/handler"
	"golang.org/x/oauth2"
)

type Config struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	RedirectURL  string `json:"redirect_url"`
	DiscordToken string `json:"discord_token"`
	JWTKey       string `json:"jwt_key"`
}

func main() {
	// read config file
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		panic("config file not found")
	}
	var config *Config
	err = json.Unmarshal(file, &config)
	if err != nil {
		panic("mangled config file, fix it")
	}

	oauthConfig := &oauth2.Config{
		ClientID:     config.ClientID,
		ClientSecret: config.ClientSecret,
		Endpoint:     oauth2.Endpoint{TokenURL: "https://discord.com/api/oauth2/token"},
		RedirectURL:  config.RedirectURL,
	}

	disc, err := discord.NewDiscordService(config.DiscordToken)
	if err != nil {
		panic(err)
	}
	disc.Open()

	jwtUtil := service.NewJWTUtil([]byte(config.JWTKey))

	conf := &handler.Config{Discord: disc, OauthConfig: oauthConfig, JwtUtil: jwtUtil}

	r := handler.NewHandler(conf)
	r.Run(":4444")
}
