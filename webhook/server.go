package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/request", getReq)
	router.POST("/webhook", write)
	router.Run(":4949")
}
