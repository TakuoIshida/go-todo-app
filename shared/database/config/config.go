package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoEnv        string
	DbHost       string
	DbUser       string
	Db           string
	DbPassword   string
	DbPort       string
	DbPortLatest string
	TZ           string
	Port         string
}

var Conf Config

func LoadConfig() {
	// .env ファイルを読み込む
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	Conf = Config{
		GoEnv:        os.Getenv("GO_ENV"),
		DbHost:       os.Getenv("POSTGRES_HOST"),
		DbUser:       os.Getenv("POSTGRES_USER"),
		Db:           os.Getenv("POSTGRES_DB"),
		DbPassword:   os.Getenv("POSTGRES_PASSWORD"),
		DbPort:       os.Getenv("POSTGRES_PORT"),
		DbPortLatest: os.Getenv("POSTGRES_PORT_LATEST"),
		TZ:           os.Getenv("TZ"),
		Port:         os.Getenv("PORT"),
	}
}
