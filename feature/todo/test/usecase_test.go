package todo_test

import (
	"errors"
	"fmt"
	"go-todo-app/feature/todo"
	"go-todo-app/feature/user"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func generateMaxLenStr() string {
	str := ""
	for i := 0; i < 256; i++ {
		str += "a"
	}
	return str
}

func TestTodoUsecaseImpl_Create(t *testing.T) {
	// Given
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
	serviceMock := &ITodoServiceMock{
		CreateFunc: func(ctx *gin.Context, userContext user.UserContext, t *todo.Todo, session *gorm.DB) error {
			// Do nothing
			return nil
		},
	}

	mockDb, mock := GetNewDbMock()

	type fields struct {
		todoService *ITodoServiceMock
		db          *gorm.DB
	}
	type args struct {
		ctx         *gin.Context
		userContext user.UserContext
		req         todo.CreateTodoRequest
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "正常: todo 作成",
			fields: fields{
				todoService: serviceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				req: todo.CreateTodoRequest{
					Title:       "Title",
					Description: "Description",
				},
			},
			want: nil,
		},
		{
			name: "異常: titleが空",
			fields: fields{
				todoService: serviceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				req: todo.CreateTodoRequest{
					Title:       "",
					Description: "Description",
				},
			},
			want: errors.New("title is required"),
		},
		{
			name: "異常: Descriptionが空",
			fields: fields{
				todoService: serviceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				req: todo.CreateTodoRequest{
					Title:       "Title",
					Description: "",
				},
			},
			want: errors.New("description is required"),
		},
		{
			name: "異常: Titleが長すぎる",
			fields: fields{
				todoService: serviceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				req: todo.CreateTodoRequest{
					Title:       generateMaxLenStr(),
					Description: "Description",
				},
			},
			want: errors.New("title is too long"),
		},
		{
			name: "異常: Descriptionが長すぎる",
			fields: fields{
				todoService: serviceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				req: todo.CreateTodoRequest{
					Title:       "Title",
					Description: generateMaxLenStr(),
				},
			},
			want: errors.New("description is too long"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Then
			usecase := todo.NewTodoUsecaseImpl(serviceMock, tt.fields.db)
			mock.ExpectBegin()
			mock.ExpectExec(fmt.Sprintf("SET app.tenant_id = '%s';", tt.args.userContext.TenantId.String())).WillReturnResult(nil)
			mock.ExpectCommit()
			got := usecase.Create(tt.args.ctx, tt.args.userContext, tt.args.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

// func TestTodoUsecaseImpl_Delete(t *testing.T) {
// 	type fields struct {
// 		todoService ITodoService
// 		db          *gorm.DB
// 	}
// 	type args struct {
// 		ctx         *gin.Context
// 		userContext user.UserContext
// 		id          uuid.UUID
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tu := &TodoUsecaseImpl{
// 				todoService: tt.fields.todoService,
// 				db:          tt.fields.db,
// 			}
// 			tu.Delete(tt.args.ctx, tt.args.userContext, tt.args.id)
// 		})
// 	}
// }

// func TestTodoUsecaseImpl_FindAll(t *testing.T) {
// 	type fields struct {
// 		todoService ITodoService
// 		db          *gorm.DB
// 	}
// 	type args struct {
// 		ctx         *gin.Context
// 		userContext user.UserContext
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   []Todo
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tu := &TodoUsecaseImpl{
// 				todoService: tt.fields.todoService,
// 				db:          tt.fields.db,
// 			}
// 			if got := tu.FindAll(tt.args.ctx, tt.args.userContext); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TodoUsecaseImpl.FindAll() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

// func TestTodoUsecaseImpl_FindById(t *testing.T) {
// 	type fields struct {
// 		todoService ITodoService
// 		db          *gorm.DB
// 	}
// 	type args struct {
// 		ctx         *gin.Context
// 		userContext user.UserContext
// 		id          uuid.UUID
// 	}
// 	tests := []struct {
// 		name   string
// 		fields fields
// 		args   args
// 		want   Todo
// 	}{
// 		// TODO: Add test cases.
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			tu := &TodoUsecaseImpl{
// 				todoService: tt.fields.todoService,
// 				db:          tt.fields.db,
// 			}
// 			if got := tu.FindById(tt.args.ctx, tt.args.userContext, tt.args.id); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("TodoUsecaseImpl.FindById() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }
