package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"Coding_Challenge/internal/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env file failed to load")
	}

	r := gin.Default()
	api.RegisterRoutes(r)

	if err := r.Run(":5000"); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}