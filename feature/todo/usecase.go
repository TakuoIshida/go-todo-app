package todo

import (
	"go-todo-app/shared/database"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoUsecaseImpl struct {
	todoService ITodoService
	db          *gorm.DB
}

func NewTodoUsecaseImpl(ts ITodoService, db *gorm.DB) ITodoUsecase {
	return &TodoUsecaseImpl{
		todoService: ts,
		db:          db,
	}
}

// Create implements todoservice.ITodoService.
func (tu *TodoUsecaseImpl) Create(ctx *gin.Context, req CreateTodoRequest) {
	new, _ := New(req.Title, req.Description, req.UserId, req.TenantId)
	tu.todoService.Create(ctx, new, tu.db)
}

// Delete implements TodoService
func (tu *TodoUsecaseImpl) Delete(ctx *gin.Context, id uuid.UUID) {
	tu.todoService.Delete(ctx, id, tu.db)
}

// FindAll implements TodoService
func (tu *TodoUsecaseImpl) FindAll(ctx *gin.Context) []Todo {
	tenantId, err := uuid.Parse("a5251b75-575d-437b-aff0-a029e509ff06")
	if err != nil {
		panic(err)
	}
	return database.TenantQuery(tu.db, tenantId, func(session *gorm.DB) []Todo {
		return tu.todoService.FindAll(ctx, session)
	})
}

// FindById implements TodoService
func (tu *TodoUsecaseImpl) FindById(ctx *gin.Context, id uuid.UUID) Todo {
	return tu.todoService.FindById(ctx, id, tu.db)
}

// // Update implements TodoService
// func (tu *TodoUsecaseImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.TodoRepository.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.TodoRepository.Update(tagData)
// }
