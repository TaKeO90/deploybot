package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	router := gin.Default()

	router.GET("/request", getReq)
	router.POST("/webhook", write)
	router.Run(":" + port)
}
