package setting

type Config struct {
	Mysql MySQLSetting `mapstructure:"mysql"`
}

type MySQLSetting struct {
	Host             string `mapstructure:"host"`
	Port             string `mapstructure:"port"`
	Username         string `mapstructure:"username"`
	Password         string `mapstructure:"password"`
	DbName           string `mapstructure:"dbname"`
	MaxIdleCons      int    `mapstructure:"maxIdleCons"`
	MaxOpenCons      int    `mapstructure:"maxOpenCons"`
	ConnnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}
