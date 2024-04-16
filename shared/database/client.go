package database

import (
	"errors"
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

func TenantTx(db *gorm.DB, tenantId uuid.UUID, callback func(session *gorm.DB) error) error {
	db.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "tenant.",
	}
	// Set the tenant ID for RLS using the context option
	return db.Transaction(func(session *gorm.DB) error {
		// escapeをするとsyntaxエラーになるため、Sprintfで対応。
		session.Exec(fmt.Sprintf("SET app.tenant_id = '%s';", tenantId.String()))
		// session.Exec("SET app.tenant_id = $1;", tenantId.String())
		err := callback(session)
		if err != nil {
			return errors.New("error in TenantTx: rollback. message: " + err.Error())
		}
		return nil
	})
}

func TenantQuery[T any](db *gorm.DB, tenantId uuid.UUID, callback func(session *gorm.DB) (T, error)) (T, error) {
	db.Config.NamingStrategy = schema.NamingStrategy{
		TablePrefix: "tenant.",
	}
	// Set the tenant ID for RLS using the context option
	var result T
	var err error
	db.Connection(func(session *gorm.DB) error {
		// escapeをするとsyntaxエラーになるため、Sprintfで対応。
		session.Exec(fmt.Sprintf("SET app.tenant_id = '%s';", tenantId.String()))
		result, err = callback(session)
		if err != nil {
			return errors.New("error in TenantQuery" + err.Error())
		}
		return err
	})

	return result, err
}
