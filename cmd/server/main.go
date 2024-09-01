package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", Pong)
		v1.PUT("/pong", Pong)
		v1.POST("/ping", Pong)
		v1.DELETE("/pong", Pong)
		v1.PATCH("/ping", Pong)
		v1.OPTIONS("/pong", Pong)
	}
	r.Run(":8002") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func Pong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping",
	})
}
