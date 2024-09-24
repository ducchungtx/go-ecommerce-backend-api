package user

import (
	"goecommerce/internal/wire"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (pr *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public route
	// non dependency injection
	// ur := repo.NewUserRepository()
	// us := service.NewUserService(ur)
	// userHandlerNonDependencyInjection := controller.NewUserController(us)

	userController, _ := wire.InitUserRouterHandler()

	// Dependency injection
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.Register)
		userRouterPublic.GET("/otp")
	}

	// private route
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/profile")
	}
}
