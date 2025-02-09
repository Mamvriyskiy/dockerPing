package main

import (
	"fmt"
	"os"
	"github.com/spf13/viper"
	"github.com/joho/godotenv"
	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/repository"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/services"
	"github.com/Mamvriyskiy/dockerPing/backend/internal/handler"
	app "github.com/Mamvriyskiy/dockerPing/backend/internal/app"
)

func main() {
	fmt.Println("Start server ...")

	if err := initConfig(); err != nil {
		logger.Log("Error", "Error configuring DB", err)
		return 
	}

	logger.Log("Info", "Configuration initialization started", nil)

	if err := godotenv.Load("configs/.env"); err != nil {
		logger.Log("Error","Error loading environment file", err)
		return
	}

	logger.Log("Info", "Loading environment configuration", nil)

	db, err := repository.NewPostgresDB(&repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logger.Log("Error", "Failed to establish connection to the database", nil)
		return
	}

	repos := repository.NewRepository(db)
	services := services.NewServicesPsql(repos)
	handlers := handler.NewHandler(services)

	logger.Log("Info", "Successfully established connection to the database", nil)


	srv := new(app.Server)
	if err := srv.Run("8000", handlers.InitRouters()); err != nil {
		logger.Log("Error", "Error occurred while starting the HTTP server", err)
		return
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")

	return viper.ReadInConfig()
}
