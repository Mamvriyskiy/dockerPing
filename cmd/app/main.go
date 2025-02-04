package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Start server ...")

	if err := initConfig(); err != nil {
		// logger.Log("Error", "initCongig", "Error config DB:", err, "")
		return 
	}

	// logger.Log(InitConfig)

	if err := godotenv.Load(); err != nil {
		// logger.Log()
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
