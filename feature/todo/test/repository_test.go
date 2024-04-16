package todo_test

import (
	"fmt"
	"go-todo-app/feature/todo"
	"go-todo-app/feature/user"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	}{
		{
			name: "正常： Todo作成成功",
			tr:   repository,
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				todo:        testTodo,
				session:     mockDb,
			},
			want: nil,
		},
		// {
		// 	name: "異常： Todo作成失敗",
		// 	tr:   repository,
		// 	args: args{
		// 		ctx:         ctx,
		// 		userContext: testUserContext,
		// 		todo:        testTodo,
		// 		session:     mockDb,
		// 	},
		// 	want: nil,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.ExpectBegin()
			mock.ExpectExec(fmt.Sprintf(`SET app.tenant_id = '%s';`, tt.args.userContext.TenantId.String())).WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			mock.ExpectCommit()

			// tt.args.session.Create(&tt.args.todo)
			tt.tr.Create(tt.args.ctx, tt.args.userContext, &tt.args.todo, tt.args.session)

			mock.ExpectExec(
				regexp.QuoteMeta(
					`SET app.tenant_id = $1;`+
						`INSERT INTO todos ("id", "tenant_id", "title", "description", "is_deleted", "created_at", "create_user_id",
				"updated_at", "update_user_id")
VALUES ($2, $3, $4, $5, $6, $7, $8, $9, $10);`)).WithArgs(
				testTodo.TenantId,
				testTodo.CreateUserId,
				testTodo.UpdateUserId,
				testTodo.CreatedAt,
				testTodo.UpdatedAt,
				testTodo.TenantId,
				testTodo.Id,
				testTodo.Title,
				testTodo.Description,
				testTodo.IsDeleted,
			).WillReturnResult(sqlmock.NewResult(0, 0))

			// assert.Equal(t, tt.want, got)
		})
	}
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
