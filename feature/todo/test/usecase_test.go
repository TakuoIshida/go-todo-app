package todo_test

import (
	"errors"
	"fmt"
	"go-todo-app/feature/todo"
	"go-todo-app/feature/user"
	"testing"
	"time"

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
			usecase := todo.NewTodoUsecaseImpl(tt.fields.todoService, tt.fields.db)
			mock.ExpectBegin()
			mock.ExpectExec(fmt.Sprintf(`SET app.tenant_id = '%s';`, tt.args.userContext.TenantId.String())).WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			// mock.ExpectExec(regexp.QuoteMeta(`SET app.tenant_id = $1;`)).WithArgs(tt.args.userContext.TenantId.String()).WillReturnResult(sqlmock.NewResult(0, 0))
			mock.ExpectCommit()
			got := usecase.Create(tt.args.ctx, tt.args.userContext, tt.args.req)
			assert.Equal(t, tt.want, got)

			// verify
			if tt.want == nil {
				assert.Equal(t, len(tt.fields.todoService.calls.Create), 1)
			}
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
			usecase := todo.NewTodoUsecaseImpl(tt.fields.todoService, tt.fields.db)
			mock.ExpectBegin()
			mock.ExpectExec(fmt.Sprintf(`SET app.tenant_id = '%s';`, tt.args.userContext.TenantId.String())).WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			// mock.ExpectExec(regexp.QuoteMeta(`SET app.tenant_id = $1;`)).WithArgs(tt.args.userContext.TenantId.String()).WillReturnResult(sqlmock.NewResult(0, 0))
			mock.ExpectCommit()
			got := usecase.Delete(tt.args.ctx, tt.args.userContext, tt.args.id)
			assert.Equal(t, tt.want, got)

			// verify
			if tt.want == nil {
				assert.Equal(t, len(tt.fields.todoService.calls.Delete), 1)
			}
		})
	}
}

func TestTodoUsecaseImpl_FindById(t *testing.T) {
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
	testTodoId := uuid.New()
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
	ctx := &gin.Context{}
	serviceMock := &ITodoServiceMock{
		FindByIdFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) (todo.Todo, error) {
			// Do nothing
			return testTodo, nil
		},
	}
	notFoundServiceMock := &ITodoServiceMock{
		FindByIdFunc: func(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) (todo.Todo, error) {
			// Do nothing
			err := errors.New("Todo not found.")
			return todo.Todo{}, err
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
		want   todo.Todo
		err    error
	}{
		{
			name: "正常: todo 取得",
			fields: fields{
				todoService: serviceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				id:          testTodoId,
			},
			want: testTodo,
			err:  nil,
		},
		{
			name: "異常: todo 取得できなかった場合",
			fields: fields{
				todoService: notFoundServiceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
				id:          uuid.New(),
			},
			want: todo.Todo{},
			err:  errors.New("Todo not found."),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Then
			usecase := todo.NewTodoUsecaseImpl(tt.fields.todoService, tt.fields.db)
			mock.ExpectExec(fmt.Sprintf(`SET app.tenant_id = '%s';`, tt.args.userContext.TenantId.String())).WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			got, err := usecase.FindById(tt.args.ctx, tt.args.userContext, tt.args.id)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, len(tt.fields.todoService.calls.FindById), 1)
		})
	}
}

func TestTodoUsecaseImpl_FindAll(t *testing.T) {
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
	testTodoId := uuid.New()
	testTodoId2 := uuid.New()
	testTodos := []todo.Todo{
		{
			Id:           testTodoId,
			TenantId:     testTenantId,
			Title:        "Title",
			Description:  "Description",
			IsDeleted:    false,
			CreatedAt:    time.Now(),
			CreateUserId: testUserId,
			UpdatedAt:    time.Now(),
			UpdateUserId: testUserId,
		},
		{
			Id:           testTodoId2,
			TenantId:     testTenantId,
			Title:        "Title2",
			Description:  "Description2",
			IsDeleted:    false,
			CreatedAt:    time.Now(),
			CreateUserId: testUserId,
			UpdatedAt:    time.Now(),
			UpdateUserId: testUserId,
		},
	}
	ctx := &gin.Context{}
	serviceMock := &ITodoServiceMock{
		FindAllFunc: func(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) ([]todo.Todo, error) {
			// Do nothing
			return testTodos, nil
		},
	}
	emptyServiceMock := &ITodoServiceMock{
		FindAllFunc: func(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) ([]todo.Todo, error) {
			// Do nothing
			return []todo.Todo{}, nil
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
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []todo.Todo
		err    error
	}{
		{
			name: "正常: todoList 取得",
			fields: fields{
				todoService: serviceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
			},
			want: testTodos,
			err:  nil,
		},
		{
			name: "正常: todoListが空の場合",
			fields: fields{
				todoService: emptyServiceMock,
				db:          mockDb,
			},
			args: args{
				ctx:         ctx,
				userContext: testUserContext,
			},
			want: []todo.Todo{},
			err:  nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Then
			usecase := todo.NewTodoUsecaseImpl(tt.fields.todoService, tt.fields.db)
			mock.ExpectExec(fmt.Sprintf(`SET app.tenant_id = '%s';`, tt.args.userContext.TenantId.String())).WithoutArgs().WillReturnResult(sqlmock.NewResult(0, 0))
			got, err := usecase.FindAll(tt.args.ctx, tt.args.userContext)
			assert.Equal(t, tt.want, got)
			assert.Equal(t, tt.err, err)
			assert.Equal(t, len(tt.fields.todoService.calls.FindAll), 1)
		})
	}
}
