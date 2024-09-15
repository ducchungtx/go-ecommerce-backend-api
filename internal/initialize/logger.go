package initialize

import (
	"goecommerce/global"
	"goecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.NewLogger(global.Config.Logger)
}
