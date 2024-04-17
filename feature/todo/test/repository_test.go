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
	testTenantId := uuid.New()
	testUserId := uuid.New()
	testTodoId := uuid.New()
	testTodoId2 := uuid.New()
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
	testDeletedTodo := todo.Todo{
		Id:           testTodoId2,
		TenantId:     testTenantId,
		Title:        "Title",
		Description:  "Description",
		IsDeleted:    true, //削除済み
		CreatedAt:    time.Now(),
		CreateUserId: testUserId,
		UpdatedAt:    time.Now(),
		UpdateUserId: testUserId,
	}
	mockDb, mock := GetNewDbMock()
	repository := todo.NewTodoRepositoryImpl()

	t.Run("正常：todoを取得できる", func(t *testing.T) {
		columns := []string{"id", "tenant_id", "title", "description", "is_deleted", "created_at", "create_user_id", "updated_at", "update_user_id"}
		mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "todos" WHERE is_deleted = $1 AND "todos"."id" = $2`)).WithArgs(
			false,
			testTodo.Id,
		).WillReturnRows(sqlmock.NewRows(columns).AddRow(testTodo.Id, testTodo.TenantId, testTodo.Title, testTodo.Description, testTodo.IsDeleted, testTodo.CreatedAt, testTodo.CreateUserId, testTodo.UpdatedAt, testTodo.UpdateUserId))

		repository.FindById(ctx, testUserContext, testTodoId, mockDb)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})

	t.Run("異常：削除されたtodoは取得できない", func(t *testing.T) {
		mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "todos" WHERE is_deleted = $1 AND "todos"."id" = $2`)).WithArgs(
			false,
			testDeletedTodo.Id,
		).WillReturnError(errors.New("fail to find todo"))

		_, err := repository.FindById(ctx, testUserContext, testDeletedTodo.Id, mockDb)

		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if assert.Error(t, err) {
			assert.Equal(t, "Todo not found. message: fail to find todo", err.Error())
		}
	})

	t.Run("異常：todoを取得でエラー", func(t *testing.T) {
		mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "todos" WHERE is_deleted = $1 AND "todos"."id" = $2`)).WithArgs(
			false,
			testTodo.Id,
		).WillReturnError(errors.New("fail to find todo"))

		_, err := repository.FindById(ctx, testUserContext, testTodoId, mockDb)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if assert.Error(t, err) {
			assert.Equal(t, "Todo not found. message: fail to find todo", err.Error())
		}
	})
}

func TestTodoRepositoryImpl_FindAll(t *testing.T) {
	testTenantId := uuid.New()
	testUserId := uuid.New()
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
		Id:           uuid.New(),
		TenantId:     testTenantId,
		Title:        "Title",
		Description:  "Description",
		IsDeleted:    false,
		CreatedAt:    time.Now(),
		CreateUserId: testUserId,
		UpdatedAt:    time.Now(),
		UpdateUserId: testUserId,
	}
	testTodo2 := todo.Todo{
		Id:           uuid.New(),
		TenantId:     testTenantId,
		Title:        "Title",
		Description:  "Description",
		IsDeleted:    false,
		CreatedAt:    time.Now(),
		CreateUserId: testUserId,
		UpdatedAt:    time.Now(),
		UpdateUserId: testUserId,
	}
	mockDb, mock := GetNewDbMock()
	repository := todo.NewTodoRepositoryImpl()

	query := `SELECT * FROM "todos" WHERE is_deleted = $1`

	t.Run("正常：todo Listを取得できる", func(t *testing.T) {
		columns := []string{"id", "tenant_id", "title", "description", "is_deleted", "created_at", "create_user_id", "updated_at", "update_user_id"}
		mock.ExpectQuery(
			regexp.QuoteMeta(query)).WithArgs(
			false,
		).WillReturnRows(sqlmock.NewRows(columns).AddRow(
			testTodo.Id, testTodo.TenantId, testTodo.Title, testTodo.Description, testTodo.IsDeleted, testTodo.CreatedAt, testTodo.CreateUserId, testTodo.UpdatedAt, testTodo.UpdateUserId,
		).AddRow(
			testTodo2.Id, testTodo2.TenantId, testTodo2.Title, testTodo2.Description, testTodo2.IsDeleted, testTodo2.CreatedAt, testTodo2.CreateUserId, testTodo2.UpdatedAt, testTodo2.UpdateUserId,
		))

		_, err := repository.FindAll(ctx, testUserContext, mockDb)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
		assert.Nil(t, err)
	})

	t.Run("異常：todoを取得でエラー", func(t *testing.T) {
		mock.ExpectQuery(
			regexp.QuoteMeta(query)).WithArgs(
			false,
		).WillReturnError(errors.New("fail to find todo list"))

		_, err := repository.FindAll(ctx, testUserContext, mockDb)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}

		if assert.Error(t, err) {
			assert.Equal(t, "Todo not found. message: fail to find todo list", err.Error())
		}
	})
}

func TestTodoRepositoryImpl_Delete(t *testing.T) {
	testTenantId := uuid.New()
	testUserId := uuid.New()
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
		Id:           uuid.New(),
		TenantId:     testTenantId,
		Title:        "Title",
		Description:  "Description",
		IsDeleted:    false,
		CreatedAt:    time.Now(),
		CreateUserId: testUserId,
		UpdatedAt:    time.Now(),
		UpdateUserId: testUserId,
	}

	mockDb, mock := GetNewDbMock()
	repository := todo.NewTodoRepositoryImpl()

	t.Run("正常：todo を削除できる", func(t *testing.T) {
		columns := []string{"id", "tenant_id", "title", "description", "is_deleted", "created_at", "create_user_id", "updated_at", "update_user_id"}
		mock.ExpectQuery(
			regexp.QuoteMeta(`SELECT * FROM "todos" WHERE is_deleted = $1 AND "todos"."id" = $2`)).WithArgs(
			false,
			testTodo.Id,
		).WillReturnRows(sqlmock.NewRows(columns).AddRow(testTodo.Id, testTodo.TenantId, testTodo.Title, testTodo.Description, testTodo.IsDeleted, testTodo.CreatedAt, testTodo.CreateUserId, testTodo.UpdatedAt, testTodo.UpdateUserId))

		mock.ExpectBegin()
		mock.ExpectExec(
			regexp.QuoteMeta(`UPDATE "todos" SET`)).WithArgs(
			true,
			sqlmock.AnyArg(),
			testUserContext.Id,
			false,
			testTodo.Id,
		).WillReturnResult(sqlmock.NewResult(0, 1))
		mock.ExpectCommit()

		repository.Delete(ctx, testUserContext, testTodo.Id, mockDb)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	})
}
