package main

import (
	"fmt"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	"github.com/Mamvriyskiy/dockerPing/tree/develop/logger"
)

func main() {
	fmt.Println("Start server ...")

	if err := initConfig(); err != nil {
		logger.Log("Error", "initCongig", "Error config DB:", err, "")
		return 
	}

	logger.Log("Info", "", "InitConfig", nil)

	if err := godotenv.Load(); err != nil {
		logger.Log("Error", "Load", "Load env file:", err, "")
		return
	}

	logger.Log("Info", "", "Load env", nil)


	// repos := repository.NewRepository(db)
	// services := service.NewServicesPsql(repos)
	// handlers := handler.NewHandler(services)

	logger.Log("Info", "", "The connection to the database is established", nil)


	srv := new(pkg.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		logger.Log("Error", "Run", "Error occurred while running http server:", err, "")
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
