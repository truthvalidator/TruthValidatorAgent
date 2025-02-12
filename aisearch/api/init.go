package main

import (
	"os"
	"ssaisearch/api/shared"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/sashabaranov/go-openai"
)

func initAll() {
	godotenv.Load(".env")
	initChatGPT()
	initGin()
}

func initChatGPT() {
	token := os.Getenv("OPENAI")
	ep := os.Getenv("OPENAI_ENDPOINT")
	cfg := openai.DefaultConfig(token)
	if ep != "" {
		cfg.BaseURL = ep
	}
	shared.Cli = openai.NewClientWithConfig(cfg)
}

var g *gin.Engine

func initGin() {
	g = gin.New()
	g.Use(gin.Logger(), gin.Recovery())

	config := cors.DefaultConfig()
	if strings.TrimSpace(os.Getenv("DEBUG")) == "1" {
		gin.SetMode(gin.DebugMode)
		config.AllowAllOrigins = true
		g.Use(cors.New(config))
	} else {
		gin.SetMode(gin.ReleaseMode)
		ao := os.Getenv("ALLOW_ORIGINS")
		config.AllowOrigins = strings.Split(ao, " ")
	}
	g.Use(cors.New(config))
}

func startGin() {
	listenAddr := os.Getenv("LISTEN_ADDR")
	if listenAddr == "" {
		listenAddr = ":8080"
	}
	g.Run(listenAddr)
}
