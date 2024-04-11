package todo

import (
	"fmt"
	"go-todo-app/helper"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TodoRepositoryImpl struct {
	Db *gorm.DB
}

func NewTodoRepositoryImpl(Db *gorm.DB) ITodoRepository {
	return &TodoRepositoryImpl{Db: Db}
}

// Delete implements TodoRepository
func (t *TodoRepositoryImpl) Delete(ctx *gin.Context, id uuid.UUID) {
	var todo Todo
	result := t.Db.Where("id = ?", id).Delete(&todo)
	if result.RowsAffected == 0 {
		helper.ErrorPanic(result.Error)
	}
	fmt.Println("deleted")
}

// FindAll implements TodoRepository
func (t *TodoRepositoryImpl) FindAll(ctx *gin.Context) []Todo {
	var todo []Todo
	result := t.Db.Find(&todo)
	if result.Error != nil {
		helper.ErrorPanic(result.Error)
	}
	return todo
}

// FindById implements TodoRepository
func (t *TodoRepositoryImpl) FindById(ctx *gin.Context, id uuid.UUID) Todo {
	// genから生成したmodelでも取得できるがmappingが大変。
	// DDDでentity = tableの場合 => GORMのdomain/modelのentityのまま利用した方が良さそう
	// DDDでentity = tableの場合 => GORMのdomain/modelのentityのまま利用した方が良さそう
	// result, err := query.Todo.Where(query.Todo.ID.Eq(id)).First()
	// if err != nil {
	// 	panic(err)
	// }
	var todo Todo
	result := t.Db.Find(&todo, id)
	fmt.Println(result)

	return todo
}

// Save implements TodoRepository
func (t *TodoRepositoryImpl) Save(ctx *gin.Context, todo *Todo) {
	result := t.Db.Create(&todo)
	helper.ErrorPanic(result.Error)
}
