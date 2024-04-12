package database

import (
	"fmt"
	"go-todo-app/shared/database/config"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBClientConnector struct {
	DB *gorm.DB
}

func NewClientConnector() *DBClientConnector {
	cfg := config.Conf
	// NOTE: db is the service name of the database in docker-compose
	var dsn string
	if cfg.GoEnv == "local" {
		dsn = fmt.Sprintf("postgres://%s:%s@db/%s", cfg.DbUser, cfg.DbPassword, cfg.Db)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.Db)
	}
	DB, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if cfg.GoEnv == "local" {
		DB.Logger.LogMode(logger.Info)
	}

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("db connected!!")

	return &DBClientConnector{
		DB: DB,
	}
}
