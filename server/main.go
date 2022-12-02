package main

import (
	"log"
	"os"

	"favorite/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Init env configuration is required
	initEnv()
	// Init Web Server
	initServer()
}

func initEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func initServer() {

	r := gin.New()
	router.Init(r)
	// router.InitWebsocket(r)
	if err := r.Run(os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT")); err != nil {
		log.Fatal("failed run app: ", err)
	}
}
