package main

import (
	"fmt"
	"go-todo-app/shared/database/config"
	"go-todo-app/shared/database/migration/tenant"
	"log"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Load the configuration
	config.LoadConfig()

	// Initialize the database
	db := dbInit()

	// Drop tables (constraints will be taken care of)
	// db.Migrator().DropTable(&tenant.Todo{}, &tenant.AttachmentFile{})

	// Perform migration using go-gormigrate and rollback if it fails
	migrate(db)

	// Rollback the last successful migration
}

func dbInit() *gorm.DB {
	cfg := config.Conf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	GenerateTableStruct(db)
	fmt.Println("db initialized!")
	return db
}

func GenerateTableStruct(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath: "query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db) // reuse your gorm db

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable(),
	)
	// Generate the code
	g.Execute()
}

func migrate(db *gorm.DB) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, []*gormigrate.Migration{
		{
			ID: "202404070000",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time
				// tx.Migrator().CreateTable(&tenant.Todo{}, &tenant.AttachmentFile{})
				// tx.Migrator().CreateIndex(&tenant.AttachmentFile{}, "todo_id")
				return tx.Migrator().AutoMigrate(tenant.Todo{}, tenant.AttachmentFile{})
			},
			Rollback: func(tx *gorm.DB) error {
				log.Println("Migration Rollback")
				return tx.Migrator().DropTable(&tenant.Todo{}, &tenant.AttachmentFile{})
			},
		},
	})

	e := m.Migrate()
	if e != nil {
		m.RollbackLast()
		log.Fatalf("Migration failed: %v", e)
		return
	}
	log.Println("Migration did run successfully")
}
