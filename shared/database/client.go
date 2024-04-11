package database

import (
	"fmt"
	"go-todo-app/shared/database/config"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClientConnector struct {
	DB *gorm.DB
}

func NewTenantClientConnector() *DBClientConnector {
	cfg := config.Conf
	// NOTE: db is the service name of the database in docker-compose
	var dsn string
	if cfg.GoEnv == "local" {
		dsn = fmt.Sprintf("%s:%s@tcp(db)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbTenant)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbTenant)
	}
	tenantDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("tenant db connected!!")

	return &DBClientConnector{
		DB: tenantDb,
	}
}

func NewCommonClientConnector() *DBClientConnector {
	cfg := config.Conf
	// NOTE: db is the service name of the database in docker-compose
	var dsn string
	if cfg.GoEnv == "local" {
		dsn = fmt.Sprintf("%s:%s@tcp(db)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbCommon)
	} else {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbCommon)
	}
	commonDb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("common db connected!!")

	return &DBClientConnector{
		DB: commonDb,
	}
}
