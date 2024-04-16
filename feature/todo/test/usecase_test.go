package todo_test

import (
	"errors"
	"fmt"
	"go-todo-app/feature/todo"
	"go-todo-app/feature/user"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
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
			mock.ExpectExec(fmt.Sprintf(`SET app.tenant_id = '%s';`, tt.args.userContext.TenantId.String())).WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			// mock.ExpectExec(regexp.QuoteMeta(`SET app.tenant_id = $1;`)).WithArgs(tt.args.userContext.TenantId.String()).WillReturnResult(sqlmock.NewResult(0, 0))
			mock.ExpectCommit()
			got := usecase.Create(tt.args.ctx, tt.args.userContext, tt.args.req)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestTodoUsecaseImpl_Delete(t *testing.T) {
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
		DeleteFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) error {
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
		id          uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   error
	}{
		{
			name: "正常: todo 削除",
			fields: fields{
				todoService: serviceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				id:          uuid.New(),
			},
			want: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Then
			usecase := todo.NewTodoUsecaseImpl(serviceMock, tt.fields.db)
			mock.ExpectBegin()
			mock.ExpectExec(fmt.Sprintf(`SET app.tenant_id = '%s';`, tt.args.userContext.TenantId.String())).WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			// mock.ExpectExec(regexp.QuoteMeta(`SET app.tenant_id = $1;`)).WithArgs(tt.args.userContext.TenantId.String()).WillReturnResult(sqlmock.NewResult(0, 0))
			mock.ExpectCommit()
			got := usecase.Delete(tt.args.ctx, tt.args.userContext, tt.args.id)
			assert.Equal(t, tt.want, got)
		})
	}
}
