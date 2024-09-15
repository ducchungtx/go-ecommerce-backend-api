package initialize

import (
	"fmt"
	"goecommerce/global"
	"goecommerce/internal/po"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func checkErrorPanics(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}

}

func InitMysql() {
	// initialize mysql
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local" // if keep the loc=Local, it will use the system's timezone
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		// default transaction is enabled (data consistency), but we don't need it to increase the performance about 30%
		SkipDefaultTransaction: false,
	})
	checkErrorPanics(err, "Failed to connect to mysql")
	global.Logger.Info("Successfully connected to mysql")
	global.Mdb = db
	// set mysql pool - open and idle connections and connection lifetime to avoid memory leaks, increase performance
	SetPool()
	// migrate tables
	migrateTables()
}

// InitMysql().SetPool()
func SetPool() {
	m := global.Config.Mysql
	// set mysql pool
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("Mysql error: ", err)
	}
	sqlDb.SetConnMaxIdleTime(time.Duration(m.MaxIdleConns))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifeTime))
}

func migrateTables() {
	// migrate tables
	err := global.Mdb.AutoMigrate(
		&po.User{},
		&po.Role{},
	)

	if err != nil {
		global.Logger.Error("Failed to migrate tables", zap.Error(err))
		panic(err)
	}
}
