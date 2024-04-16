package todo_test

import (
	"errors"
	"go-todo-app/feature/todo"
	"go-todo-app/feature/user"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestTodoRepositoryImpl_Create(t *testing.T) {
	// Given
	mockDb, mock := GetNewDbMock()
	testTenantId := uuid.New()
	testUserId := uuid.New()
	testTodoId := uuid.New()
	testUserContext := user.UserContext{
		Id:        testUserId,
		TenantId:  testTenantId,
		Email:     "example@gmail.com",
		LastName:  "LastName",
		FirstName: "FirstName",
		AccountId: uuid.MustParse("00000000-0000-0000-0000-000000000000"),
	}
	ctx := &gin.Context{}

	testTodo := todo.Todo{
		Id:           testTodoId,
		TenantId:     testTenantId,
		Title:        "Title",
		Description:  "Description",
		IsDeleted:    false,
		CreatedAt:    time.Now(),
		CreateUserId: testUserId,
		UpdatedAt:    time.Now(),
		UpdateUserId: testUserId,
	}
	repository := todo.NewTodoRepositoryImpl()

	t.Run("正常： Todo作成成功", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "todos"`)).WithArgs(
			testTodo.Id,
			testTodo.TenantId,
			testTodo.Title,
			testTodo.Description,
			testTodo.IsDeleted,
			testTodo.CreatedAt,
			testTodo.CreateUserId,
			testTodo.UpdatedAt,
			testTodo.UpdateUserId,
		).WillReturnResult(sqlmock.NewResult(0, 0))
		mock.ExpectCommit()

		got := repository.Create(ctx, testUserContext, &testTodo, mockDb)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		assert.Nil(t, got)
	})

	t.Run("異常： Todo作成失敗", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(
			regexp.QuoteMeta(`INSERT INTO "todos"`)).WithArgs(
			testTodo.Id,
			testTodo.TenantId,
			testTodo.Title,
			testTodo.Description,
			testTodo.IsDeleted,
			testTodo.CreatedAt,
			testTodo.CreateUserId,
			testTodo.UpdatedAt,
			testTodo.UpdateUserId,
		).WillReturnError(errors.New("fail to create todo"))
		mock.ExpectRollback()

		got := repository.Create(ctx, testUserContext, &testTodo, mockDb)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if assert.Error(t, got) {
			assert.Equal(t, "fail to create todo", got.Error())
		}
	})
}

func TestTodoRepositoryImpl_FindById(t *testing.T) {
	type args struct {
		ctx         *gin.Context
		userContext user.UserContext
		todo        todo.Todo
		session     *gorm.DB
	}
	tests := []struct {
		name string
		tr   todo.ITodoRepository
		args args
		want error
	}{{
		// TODO: Add test cases.
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func TestTodoRepositoryImpl_FindAll(t *testing.T) {
	type args struct {
		ctx         *gin.Context
		userContext user.UserContext
		todo        todo.Todo
		session     *gorm.DB
	}
	tests := []struct {
		name string
		tr   todo.ITodoRepository
		args args
		want error
	}{{
		// TODO: Add test cases.
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}

func TestTodoRepositoryImpl_Delete(t *testing.T) {
	type args struct {
		ctx         *gin.Context
		userContext user.UserContext
		todo        todo.Todo
		session     *gorm.DB
	}
	tests := []struct {
		name string
		tr   todo.ITodoRepository
		args args
		want error
	}{{
		// TODO: Add test cases.
	}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

		})
	}
}
