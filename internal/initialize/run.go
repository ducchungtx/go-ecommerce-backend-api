package initialize

import (
	"fmt"
	"goecommerce/global"

	"go.uber.org/zap"
)

func Run() {
	// load configurations
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configurations mysql", m.Username, m.Password)

	InitLogger()
	global.Logger.Info("Config loaded", zap.String("ok", "success"))
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
