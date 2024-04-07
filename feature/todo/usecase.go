package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoUsecaseImpl struct {
	TodoService ITodoService
}

func NewTodoUsecaseImpl(ts ITodoService) ITodoUsecase {
	return &TodoUsecaseImpl{
		TodoService: ts,
	}
}

// Create implements todoservice.ITodoService.
func (tu *TodoUsecaseImpl) Create(ctx *gin.Context, req CreateTodoRequest) {
	new, _ := New(req.Title, req.Description, req.UserId, req.TenantId)
	tu.TodoService.Create(ctx, new)
}

// Delete implements TodoService
func (tu *TodoUsecaseImpl) Delete(ctx *gin.Context, id uuid.UUID) {
	tu.TodoService.Delete(ctx, id)
}

// FindAll implements TodoService
func (tu *TodoUsecaseImpl) FindAll(ctx *gin.Context) []Todo {
	return tu.TodoService.FindAll(ctx)
}

// FindById implements TodoService
func (tu *TodoUsecaseImpl) FindById(ctx *gin.Context, id uuid.UUID) Todo {
	return tu.TodoService.FindById(ctx, id)
}

// // Update implements TodoService
// func (tu *TodoUsecaseImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.TodoRepository.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.TodoRepository.Update(tagData)
// }
