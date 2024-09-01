package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
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
	return r
}

func Pong(c *gin.Context) {
	name := c.DefaultQuery("name", "guest")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping" + name,
		"uid":     uid,
		"users":   []string{"chung", "kim", "park"},
	})
}
