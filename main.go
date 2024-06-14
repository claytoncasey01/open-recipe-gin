package main

import (
	"log"

	"github.com/claytoncasey01/open-recipe-gin/config"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	database.ConnectDB()

	r := gin.Default()

	r.Run() // listen and serve on 0.0.0.0:8080
	log.Println("Server started on :8080")
}