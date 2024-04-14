package todo

import (
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
func (t *TodoServiceImpl) Create(ctx *gin.Context, todo *Todo, db *gorm.DB) {
	t.TodoRepository.Save(ctx, todo, db)
}

// Delete implements TodoService
func (t *TodoServiceImpl) Delete(ctx *gin.Context, id uuid.UUID, db *gorm.DB) {
	t.TodoRepository.Delete(ctx, id, db)
}

// FindAll implements TodoService
func (t *TodoServiceImpl) FindAll(ctx *gin.Context, db *gorm.DB) []Todo {
	return t.TodoRepository.FindAll(ctx, db)
}

// FindById implements TodoService
func (t *TodoServiceImpl) FindById(ctx *gin.Context, id uuid.UUID, db *gorm.DB) Todo {
	return t.TodoRepository.FindById(ctx, id, db)
}

// // Update implements TodoService
// func (t *TodoServiceImpl) Update(ctx *gin.Context, Todo request.UpdateTodoRequest) {
// 	tagData, err := t.FindById(Todo.Id)
// 	helper.ErrorPanic(err)
// 	tagData.Name = Todo.Name
// 	t.Update(tagData)
// }
