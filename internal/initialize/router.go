package initialize

import (
	"fmt"
	c "goecommerce/internal/controller"
	"goecommerce/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func AA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> AA")
		c.Next()
		fmt.Println("After --> AA")
	}
}

func BB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before --> BB")
		c.Next()
		fmt.Println("After --> BB")
	}
}

func CC(c *gin.Context) {
	fmt.Println("Before --> CC")
	c.Next()
	fmt.Println("After --> CC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware(), BB(), CC)

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
