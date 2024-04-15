package todo_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetNewDbMock() (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, _ := sqlmock.New()
	mockDB, _ := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}))
	return mockDB, mock
}
