package global

import (
	"goecommerce/pkg/logger"
	"goecommerce/setting"

	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Mdb    *gorm.DB
)

/*
Config
Redis
Mysql
...
*/
