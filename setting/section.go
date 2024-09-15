package setting

type Config struct {
	Mysql  MySQLSetting  `mapstructure:"mysql"`
	Logger LoggerSetting `mapstructure:"logger"`
}

type MySQLSetting struct {
	Host            string `mapstructure:"host"`
	Port            int    `mapstructure:"port"`
	Username        string `mapstructure:"username"`
	Password        string `mapstructure:"password"`
	Dbname          string `mapstructure:"dbname"`
	MaxIdleConns    int    `mapstructure:"maxIdleConns"`
	MaxOpenConns    int    `mapstructure:"maxOpenConns"`
	ConnMaxLifeTime int    `mapstructure:"connMaxLifeTime"`
}

type LoggerSetting struct {
	File_log_name string `mapstructure:"file_log_name"`
	Log_level     string `mapstructure:"log_level"`
	Max_size      int    `mapstructure:"max_size"`
	Max_age       int    `mapstructure:"max_age"`
	Max_backups   int    `mapstructure:"max_back"`
	Compress      bool   `mapstructure:"compress"`
}
