package routers

import (
	c "goecommerce/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/2024")
	{
		v1.GET("/ping", c.NewPongController().Pong)
		v1.GET("/user/1", c.NewUserController().GetUserById)
		// v1.PUT("/pong", Pong)
		// v1.POST("/ping", Pong)
		// v1.DELETE("/pong", Pong)
		// v1.PATCH("/ping", Pong)
		// v1.OPTIONS("/pong", Pong)
	}
	return r
}
