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
	tenantDBInit()
	// Drop tables (constraints will be taken care of)
	// db.Migrator().DropTable(&tenant.Todo{}, &tenant.AttachmentFile{})
	// db.Migrator().DropTable(&tenant.Todo{})

	// Perform migration using go-gormigrate and rollback if it fails

	// generateTableStruct(db)

	// Rollback the last successful migration
}

func tenantDBInit() {
	cfg := config.Conf
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", cfg.DbUser, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbTenant)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}

	fmt.Println("tenant db connected!")

	// migrate tenant db
	tenantMigrations := []*gormigrate.Migration{
		{
			ID: "202404070000",
			Migrate: func(tx *gorm.DB) error {
				// it's a good pratice to copy the struct inside the function,
				// so side effects are prevented if the original struct changes during the time
				return tx.Migrator().AutoMigrate(&tenant.Todo{}, &tenant.AttachmentFile{})
			},
			Rollback: func(tx *gorm.DB) error {
				log.Println("Migration Rollback")
				return tx.Migrator().DropTable(&tenant.Todo{}, &tenant.AttachmentFile{})
			},
		},
	}
	migrate(db, tenantMigrations)
	fmt.Println("tenant db migrated!")

	generateTableStruct(db)
	fmt.Println("tenant struct generated!")

}

func generateTableStruct(db *gorm.DB) {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "shared/database",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
		ModelPkgPath: "./tenant/model",
	})

	g.UseDB(db) // reuse your gorm db

	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}

func migrate(db *gorm.DB, migrations []*gormigrate.Migration) {
	m := gormigrate.New(db, gormigrate.DefaultOptions, migrations)

	e := m.Migrate()
	if e != nil {
		m.RollbackLast()
		log.Fatalf("Migration failed: %v", e)
		return
	}
	log.Println("Migration did run successfully")
}
