package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Databases []struct {
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
	} `mapstructure:"databases"`
}

func main() {
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

	// read databases
	var config Config
	// read config into struct
	if err := viper.Unmarshal(&config); err != nil {
		panic("Failed to unmarshal config: " + err.Error())
	}

	// print config
	fmt.Println("server.port:", config.Server.Port) // server.port: 8080

	for _, db := range config.Databases {
		// fmt.Println("database:", db)
		fmt.Printf("database: user=%s password=%s host=%s\n", db.User, db.Password, db.Host)
	}
}
