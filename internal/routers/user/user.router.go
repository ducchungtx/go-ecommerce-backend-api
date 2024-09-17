package user

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public route
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register")
		userRouterPublic.GET("/otp")
	}

	// private route
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/profile")
	}
}
