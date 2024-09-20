package initialize

import (
	"goecommerce/global"
	"goecommerce/internal/routers"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}
	// middleware
	// r.Use() // logging
	// r.Use() // cross
	// r.Use() // limiter global

	mangagerRouter := routers.RouterGroupApp.Manager
	userRouter := routers.RouterGroupApp.User

	MainGroup := r.Group("/api/2024")
	{
		MainGroup.GET("/checkStatus") // check status
	}
	{
		userRouter.InitUserRouter(MainGroup)    // user router
		userRouter.InitProductRouter(MainGroup) // product router
	}
	{
		mangagerRouter.InitUserRouter(MainGroup)  // manager router
		mangagerRouter.InitAdminRouter(MainGroup) // admin router
	}

	return r
}
