package database

import (
	"fmt"
	"go-todo-app/shared/database/config"
	"log"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClientConnector struct {
	DB *gorm.DB
}

func NewDBClientConnector() *DBClientConnector {
	if config.Conf.GIN_MODE == gin.ReleaseMode || config.Conf.GIN_MODE == gin.TestMode {
		db, err := connectWithCloudSql()
		if err != nil {
			log.Fatalf("cannot connect with cloud db")
		}
		return &DBClientConnector{
			DB: db,
		}
	} else {
		// local: config.Conf.GIN_MODE == gin.DebugMode
		db, err := connectWithLocalDB()
		if err != nil {
			log.Fatalf("cannot connect with local db")
		}
		return &DBClientConnector{
			DB: db,
		}
	}
}

func connectWithLocalDB() (*gorm.DB, error) {
	fmt.Println("connectWithLocalDB")
	cfg := config.Conf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn))

	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	fmt.Println("db connected!!")
	// query.SetDefault(db)
	// fmt.Println("SetDefault!!")
	return db, err
}

func connectWithCloudSql() (*gorm.DB, error) {
	fmt.Println("connectWithCloudSql")
	cfg := config.Conf
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", cfg.DbHost, cfg.DbUser, cfg.DbPassword, cfg.DbName, cfg.DbPort)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		DriverName: "cloudsqlpostgres",
		DSN:        dsn,
	}))
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	return gormDB, nil
}
