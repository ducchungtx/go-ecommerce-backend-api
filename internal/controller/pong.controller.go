package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PongController struct{}

func NewPongController() *PongController {
	return &PongController{}
}

func (p *PongController) Pong(c *gin.Context) {
	fmt.Println("--> PongController.Pong")
	name := c.DefaultQuery("name", "guest")
	uid := c.Query("uid")
	c.JSON(http.StatusOK, gin.H{
		"message": "pong...ping" + name,
		"uid":     uid,
		"users":   []string{"chung", "kim", "park"},
	})
}
