package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoEnv      string
	DbHost     string
	DbUser     string
	DbCommon   string
	DbTenant   string
	DbPassword string
	DbPort     string
	TZ         string
}

var Conf Config

func LoadConfig() {
	// .env ファイルを読み込む
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	Conf = Config{
		GoEnv:      os.Getenv("GO_ENV"),
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbCommon:   os.Getenv("DB_COMMON"),
		DbTenant:   os.Getenv("DB_TENANT"),
		DbPassword: os.Getenv("DB_PASS"),
		DbPort:     os.Getenv("DB_PORT"),
		TZ:         os.Getenv("TZ"),
	}
}
