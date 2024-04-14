package database

import (
	"fmt"
	"go-todo-app/shared/database/config"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/mysql"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DBClientConnector struct {
	DB *gorm.DB
}

func NewClientConnector() *DBClientConnector {
	cfg := config.Conf
	// NOTE: db is the service name of the database in docker-compose
	var dsn string
	if cfg.GoEnv == "local" {
		dsn = fmt.Sprintf("postgres://%s:%s@db/%s", cfg.DbTenantUser, cfg.DbTenantUserPass, cfg.Db)
	} else {
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s", cfg.DbTenantUser, cfg.DbTenantUserPass, cfg.DbHost, cfg.DbPort, cfg.Db)
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

func TenantTx(db *gorm.DB, tenantId string, callback func() error) error {
	db.Set("app.tenant_id", tenantId)
	db.Transaction(func(tx *gorm.DB) error {
		if err := callback(); err != nil {
			return err
		}
		return nil
	})
	db.Set("app.tenant_id", "")

	return nil
}

func TenantQuery[T interface{}](db *gorm.DB, tenantId uuid.UUID, callback func(session *gorm.DB) T) T {
	db.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "tenant.",
	}
	// Set the tenant ID for RLS using the context option
	var result T
	db.Connection(func(session *gorm.DB) error {
		// escapeをするとsyntaxエラーになるため、Sprintfで対応。
		session.Exec(fmt.Sprintf("SET app.tenant_id = '%s';", tenantId.String()))
		result = callback(session)
		return nil
	})

	return result
}

// func getSearchPath(db *gorm.DB) {
// 	var currentSchema string
// 	rows, err := db.Debug().Raw("show search_path").Rows()
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	rows.Next() // to get the first and only result
// 	rows.Scan(&currentSchema)
// 	fmt.Printf("Search Path: %v\n", currentSchema)

// }
