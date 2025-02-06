package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/internal/repository"
	"github.com/Mamvriyskiy/dockerPing/internal/services"
	"github.com/Mamvriyskiy/dockerPing/internal/handler"
	app "github.com/Mamvriyskiy/dockerPing/internal/app"
)

func main() {
	fmt.Println("Start server ...")

	if err := initConfig(); err != nil {
		logger.Log("Error", "initCongig", "Error config DB:", err, "")
		return 
	}

	logger.Log("Info", "", "InitConfig", nil)

	if err := godotenv.Load("configs/.env"); err != nil {
		logger.Log("Error", "Load", "Load env file:", err, "")
		return
	}

	logger.Log("Info", "", "Load env", nil)

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		// log
		return
	}

	repos := repository.NewRepository(db)
	services := services.NewServicesPsql(repos)
	handlers := handler.NewHandler(services)

	logger.Log("Info", "", "The connection to the database is established", nil)


	srv := new(app.Server)
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
