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
func (tu *TodoUsecaseImpl) Create(ctx *gin.Context, userContext user.UserContext, req CreateTodoRequest) error {
	new, err := New(req.Title, req.Description, userContext.Id, userContext.TenantId)
	if err != nil {
		return err
	}

	return database.TenantTx(tu.db, userContext.TenantId, func(session *gorm.DB) error {
		return tu.todoService.Create(ctx, userContext, new, session)
	})
}

// Delete implements TodoService
func (tu *TodoUsecaseImpl) Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) error {
	return database.TenantTx(tu.db, userContext.TenantId, func(session *gorm.DB) error {
		return tu.todoService.Delete(ctx, userContext, id, session)
	})
}

// FindAll implements TodoService
func (tu *TodoUsecaseImpl) FindAll(ctx *gin.Context, userContext user.UserContext) ([]Todo, error) {
	return database.TenantQuery(tu.db, userContext.TenantId, func(session *gorm.DB) ([]Todo, error) {
		return tu.todoService.FindAll(ctx, userContext, session)
	})
}

// FindById implements TodoService
func (tu *TodoUsecaseImpl) FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) (Todo, error) {

	return database.TenantQuery(tu.db, userContext.TenantId, func(session *gorm.DB) (Todo, error) {
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
