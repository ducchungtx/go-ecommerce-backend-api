package global

import (
	"goecommerce/pkg/logger"
	"goecommerce/setting"

	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

var (
	Config setting.Config
	Logger *logger.LoggerZap
	Rdb    *redis.Client
	Mdb    *gorm.DB
)

/*
Config
Redis
Mysql
...
*/
