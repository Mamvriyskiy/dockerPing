package repository

import (
	"fmt"

	"github.com/Mamvriyskiy/dockerPing/logger"
	"github.com/jmoiron/sqlx"

	// Импорт драйвера PostgreSQL для его регистрации
	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg *Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBName, cfg.Password, cfg.SSLMode))
	if err != nil {
		logger.Log("Error", "Failed to connect to the database", err,
			fmt.Sprintf("DB_HOST = %s, DB_NAME = %s", cfg.Host, cfg.DBName))
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		logger.Log("Error", "Error checking database connection", err)
		return nil, err
	}

	return db, nil
}
