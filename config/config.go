package config

import (
	"os"

	"github.com/joho/godotenv"
)

type AppConfig struct {
	DBConnectionString string
	DBHost             string
	DBPort             string
	DBUsername         string
	DBPassword         string
	DBName             string
	JWTSecret          string
	// Add other configuration parameters as needed
}

var AppConf AppConfig

func LoadConfig() error {
	err := godotenv.Load(".env")
	if err != nil {
		return err
	}

	AppConf.DBHost = os.Getenv("DB_HOST")
	AppConf.DBPort = os.Getenv("DB_PORT")
	AppConf.DBUsername = os.Getenv("DB_USERNAME")
	AppConf.DBPassword = os.Getenv("DB_PASSWORD")
	AppConf.DBName = os.Getenv("DB_NAME")
	AppConf.JWTSecret = os.Getenv("JWT_SECRET")
	AppConf.DBConnectionString = ""

	return nil
}
