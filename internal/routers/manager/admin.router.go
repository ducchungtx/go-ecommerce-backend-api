package manager

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (pr *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public route
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}

	// private route
	adminRouterPrivate := Router.Group("/admin/user")
	// adminRouterPrivate.Use(limiter())
	// adminRouterPrivate.Use(Authen())
	// adminRouterPrivate.Use(Permission())

	{
		adminRouterPrivate.GET("/active_user")

	}
}
