package todo

import (
	"go-todo-app/feature/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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
func (t *TodoServiceImpl) Create(ctx *gin.Context, userContext user.UserContext, todo *Todo, session *gorm.DB) error {
	return t.TodoRepository.Create(ctx, userContext, todo, session)
}

// Delete implements TodoService
func (t *TodoServiceImpl) Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) {
	t.TodoRepository.Delete(ctx, userContext, id, session)
}

// FindAll implements TodoService
func (t *TodoServiceImpl) FindAll(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []Todo {
	return t.TodoRepository.FindAll(ctx, userContext, session)
}

// FindById implements TodoService
func (t *TodoServiceImpl) FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) Todo {
	return t.TodoRepository.FindById(ctx, userContext, id, session)
}

// // Update implements TodoService
// func (t *TodoServiceImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.Update(tagData)
// }
