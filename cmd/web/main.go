package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/welcome", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "Greetings",
		})
	})

	r.Run("localhost:3000")
}
