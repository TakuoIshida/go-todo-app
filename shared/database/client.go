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
	// dsn := fmt.Sprintf("%s:%s@tcp(db)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbTenant)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbTenant)
	tenantDb, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("tenant db connected!!")
	// tenant.SetDefault(tenantDb)

	return &DBClientConnector{
		DB: tenantDb,
	}
}

func NewCommonClientConnector() *DBClientConnector {
	cfg := config.Conf
	// NOTE: db is the service name of the database in docker-compose
	// dsn := fmt.Sprintf("%s:%s@tcp(db)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbCommon)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbCommon)
	commonDb, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("common db connected!!")
	// common.SetDefault(commonDb)

	return &DBClientConnector{
		DB: commonDb,
	}
}
