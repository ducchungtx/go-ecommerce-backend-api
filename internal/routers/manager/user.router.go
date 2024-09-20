package manager

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public route
	// userRouterPublic := Router.Group("/admin/user")
	// {
	// 	userRouterPublic.POST("/register")
	// 	userRouterPublic.GET("/otp")
	// }

	// private route
	userRouterPrivate := Router.Group("/admin/user")
	// userRouterPrivate.Use(limiter())
	// userRouterPrivate.Use(Authen())
	// userRouterPrivate.Use(Permission())
	{
		userRouterPrivate.GET("/active_user")

	}
}
