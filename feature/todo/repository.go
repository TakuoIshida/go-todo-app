package todo

import (
	"go-todo-app/feature/user"
	"go-todo-app/helper"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepositoryImpl() ITodoRepository {
	return &TodoRepositoryImpl{}
}

// Save implements TodoRepository
func (t *TodoRepositoryImpl) Save(ctx *gin.Context, todo *Todo, session *gorm.DB) {
	result := session.Create(&todo)
	helper.ErrorPanic(result.Error)
}

// Delete implements TodoRepository
func (t *TodoRepositoryImpl) Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) {
	var todo Todo
	findResult := session.Find(&todo, id)
	helper.ErrorPanic(findResult.Error)

	result := session.Model(&todo).Where("is_deleted = ?", false).Updates(Todo{IsDeleted: true, UpdatedAt: time.Now(), UpdateUserId: userContext.Id})
	helper.ErrorPanic(result.Error)
}

// FindAll implements TodoRepository
func (t *TodoRepositoryImpl) FindAll(ctx *gin.Context, session *gorm.DB) []Todo {
	var todo []Todo
	result := session.Find(&todo)
	helper.ErrorPanic(result.Error)
	return todo
}

// FindById implements TodoRepository
func (t *TodoRepositoryImpl) FindById(ctx *gin.Context, id uuid.UUID, session *gorm.DB) Todo {
	var todo Todo
	result := session.Find(&todo, id)
	helper.ErrorPanic(result.Error)

	return todo
}
