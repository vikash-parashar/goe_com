package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBConnectionString string
	JWTSecret          string
	// Add other configuration parameters as needed
}

var AppConf AppConfig

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	AppConf.DBConnectionString = os.Getenv("DB_CONNECTION_STRING")
	AppConf.JWTSecret = os.Getenv("JWT_SECRET")
	// Load other config parameters

	return nil
}
