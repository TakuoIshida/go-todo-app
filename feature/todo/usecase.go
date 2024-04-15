package todo

import (
	"go-todo-app/feature/user"
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
func (tu *TodoUsecaseImpl) Create(ctx *gin.Context, userContext user.UserContext, req CreateTodoRequest) {
	new, _ := New(req.Title, req.Description, userContext.Id, userContext.TenantId)
	database.TenantTx(tu.db, userContext.TenantId, func(session *gorm.DB) error {
		tu.todoService.Create(ctx, userContext, new, session)
		return nil
	})
}

// Delete implements TodoService
func (tu *TodoUsecaseImpl) Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) {
	database.TenantTx(tu.db, userContext.TenantId, func(session *gorm.DB) error {
		tu.todoService.Delete(ctx, userContext, id, session)
		return nil
	})
}

// FindAll implements TodoService
func (tu *TodoUsecaseImpl) FindAll(ctx *gin.Context, userContext user.UserContext) []Todo {
	return database.TenantQuery(tu.db, userContext.TenantId, func(session *gorm.DB) []Todo {
		return tu.todoService.FindAll(ctx, userContext, session)
	})
}

// FindById implements TodoService
func (tu *TodoUsecaseImpl) FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) Todo {

	return database.TenantQuery(tu.db, userContext.TenantId, func(session *gorm.DB) Todo {
		return tu.todoService.FindById(ctx, userContext, id, session)
	})
}

// // Update implements TodoService
// func (tu *TodoUsecaseImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.TodoRepository.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.TodoRepository.Update(tagData)
// }
