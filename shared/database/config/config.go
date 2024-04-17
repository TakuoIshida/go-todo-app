package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	GoEnv            string
	DbHost           string
	DbUser           string
	DbTenantUser     string
	DbTenantUserPass string
	DbCommonUser     string
	DbCommonUserPass string
	Db               string
	DbPassword       string
	DbPort           string
	Port             string
}

var Conf Config

func LoadConfig() {
	// .env ファイルを読み込む
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	Conf = Config{
		GoEnv:            os.Getenv("GO_ENV"),
		DbHost:           os.Getenv("POSTGRES_HOST"),
		DbUser:           os.Getenv("POSTGRES_USER"),
		Db:               os.Getenv("POSTGRES_DB"),
		DbPassword:       os.Getenv("POSTGRES_PASSWORD"),
		DbPort:           os.Getenv("POSTGRES_PORT"),
		DbTenantUser:     os.Getenv("TENANT_USER"),
		DbTenantUserPass: os.Getenv("TENANT_USER_PASS"),
		DbCommonUser:     os.Getenv("COMMON_USER"),
		DbCommonUserPass: os.Getenv("COMMON_USER_PASS"),
		Port:             os.Getenv("PORT"),
	}
}
