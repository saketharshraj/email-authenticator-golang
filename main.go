package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)




func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading in env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "4040"
	}

	router := gin.New()
	router.Use(gin.Logger())


	router.Run(":" + port)
}