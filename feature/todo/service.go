package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoServiceImpl struct {
	TodoRepository ITodoRepository
}

func NewTodoServiceImpl(tr ITodoRepository) ITodoService {
	return &TodoServiceImpl{
		TodoRepository: tr,
	}
}

// Create implements TodoService
func (t *TodoServiceImpl) Create(ctx *gin.Context, todo *Todo) {
	t.TodoRepository.Save(ctx, todo)
}

// Delete implements TodoService
func (t *TodoServiceImpl) Delete(ctx *gin.Context, id uuid.UUID) {
	t.TodoRepository.Delete(ctx, id)
}

// FindAll implements TodoService
func (t *TodoServiceImpl) FindAll(ctx *gin.Context) []Todo {
	return t.TodoRepository.FindAll(ctx)
}

// FindById implements TodoService
func (t *TodoServiceImpl) FindById(ctx *gin.Context, id uuid.UUID) Todo {
	return t.TodoRepository.FindById(ctx, id)
}

// // Update implements TodoService
// func (t *TodoServiceImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.Update(tagData)
// }
