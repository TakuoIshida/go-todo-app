package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoEnv      string
	DbHost     string
	DbUser     string
	DbName     string
	DbPassword string
	DbPort     string
	TZ         string
}

var Conf Config

func LoadConfig() {
	// .env ファイルを読み込む
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "local")
	}
	err := godotenv.Load(fmt.Sprintf("env/%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	Conf = Config{
		GoEnv:      os.Getenv("GO_ENV"),
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbName:     os.Getenv("DB_NAME"),
		DbPassword: os.Getenv("DB_PASS"),
		DbPort:     os.Getenv("DB_PORT"),
		TZ:         os.Getenv("TZ"),
	}
}
