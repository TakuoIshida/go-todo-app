package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GIN_MODE   string
	DbHost     string
	DbUser     string
	DbName     string
	DbPassword string
	DbPort     string
	TZ         string
}

var Conf Config

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		// TODO localのみ。dev, productionでは、読み込まない
		log.Fatalf("Error loading .env file: %v", err)
	}

	Conf = Config{
		GIN_MODE:   os.Getenv("GIN_MODE"),
		DbHost:     os.Getenv("DB_HOST"),
		DbUser:     os.Getenv("DB_USER"),
		DbName:     os.Getenv("DB_NAME"),
		DbPassword: os.Getenv("DB_PASS"),
		DbPort:     os.Getenv("DB_PORT"),
		TZ:         os.Getenv("TZ"),
	}
	println(Conf.DbHost)
	println(Conf.DbUser)
	println(Conf.DbName)
}
