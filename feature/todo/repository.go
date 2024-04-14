package todo

import (
	"fmt"
	"go-todo-app/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
}

func NewTodoRepositoryImpl() ITodoRepository {
	return &TodoRepositoryImpl{}
}

// Delete implements TodoRepository
func (t *TodoRepositoryImpl) Delete(ctx *gin.Context, id uuid.UUID, session *gorm.DB) {
	var todo Todo
	result := session.Where("id = ?", id).Delete(&todo)
	if result.RowsAffected == 0 {
		helper.ErrorPanic(result.Error)
	}
	fmt.Println("deleted")
}

// FindAll implements TodoRepository
func (t *TodoRepositoryImpl) FindAll(ctx *gin.Context, session *gorm.DB) []Todo {
	var todo []Todo
	result := session.Find(&todo)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
	}
	return todo
}

// FindById implements TodoRepository
func (t *TodoRepositoryImpl) FindById(ctx *gin.Context, id uuid.UUID, session *gorm.DB) Todo {
	// genから生成したmodelでも取得できるがmappingが大変。
	// DDDでentity = tableの場合 => GORMのdomain/modelのentityのまま利用した方が良さそう
	// DDDでentity = tableの場合 => GORMのdomain/modelのentityのまま利用した方が良さそう
	// result, err := query.Todo.Where(query.Todo.ID.Eq(id)).First()
	// if err != nil {
	// 	panic(err)
	// }
	var todo Todo
	result := session.Find(&todo, id)
	fmt.Println(result)

	return todo
}

// Save implements TodoRepository
func (t *TodoRepositoryImpl) Save(ctx *gin.Context, todo *Todo, session *gorm.DB) {
	result := session.Create(&todo)
	helper.ErrorPanic(result.Error)
}
