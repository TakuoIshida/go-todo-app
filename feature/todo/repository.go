package todo

import (
	"fmt"
	"go-todo-app/feature/user"
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

// Create implements TodoRepository
func (t *TodoRepositoryImpl) Create(ctx *gin.Context, userContext user.UserContext, todo *Todo, session *gorm.DB) error {
	result := session.Create(&todo)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// FindById implements TodoRepository
func (t *TodoRepositoryImpl) FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) (Todo, error) {
	var todo Todo
	result := session.Where("is_deleted = ?", false).Find(&todo, id)
	if result.Error != nil {
		return Todo{}, fmt.Errorf("Todo not found. message: %s", result.Error)
	}
	return todo, nil
}

// FindAll implements TodoRepository
func (t *TodoRepositoryImpl) FindAll(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) ([]Todo, error) {
	var todo []Todo
	result := session.Where("is_deleted = ?", false).Find(&todo)
	if result.Error != nil {
		return nil, fmt.Errorf("Todo not found. message: %s", result.Error)
	}
	return todo, nil
}

// Delete implements TodoRepository
func (t *TodoRepositoryImpl) Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) error {
	var todo Todo
	result := session.Where("is_deleted = ?", false).Find(&todo, id)
	if result.Error != nil {
		return fmt.Errorf("Todo not found. message: %s", result.Error)
	}

	deleteResult := session.Model(&todo).Where("is_deleted = ?", false).Updates(Todo{IsDeleted: true, UpdatedAt: time.Now(), UpdateUserId: userContext.Id})
	if deleteResult.Error != nil {
		return fmt.Errorf("fail to update. id: %s, message: %s", id, deleteResult.Error)
	}
	return nil
}
