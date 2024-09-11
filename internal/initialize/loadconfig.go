package initialize

import (
	"fmt"
	"goecommerce/global"

	"github.com/spf13/viper"
)

func LoadConfig() {
	// load configurations
	viper := viper.New()
	viper.AddConfigPath("./config/") // config file path
	viper.SetConfigName("local")     // config file name without extension
	viper.SetConfigType("yaml")      // config file type

	// Read config file
	err := viper.ReadInConfig()
	if err != nil {
		panic("Failed to read config file: " + err.Error()) // handle errors reading the config file and stop the program
	}

	// read config
	// server port
	fmt.Println("server.port:", viper.GetInt("server.port")) // server.port: 8080
	// read jwt key
	fmt.Println("jwt.key:", viper.GetString("security.jwt.key")) // jwt.key: secret

	// read config into struct
	if err := viper.Unmarshal(&global.Config); err != nil {
		panic("Failed to unmarshal config: " + err.Error())
	}
}
