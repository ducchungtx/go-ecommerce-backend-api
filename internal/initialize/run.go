package initialize

import (
	"fmt"
	"goecommerce/global"
)

func Run() {
	// load configurations
	LoadConfig()
	m := global.Config.Mysql
	fmt.Println("Loading configurations mysql", m.Username, m.Password)

	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
