package todo

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TodoControllerImpl struct {
	TodoUsecase ITodoUsecase
}

func NewTodoController(tu ITodoUsecase) *TodoControllerImpl {
	return &TodoControllerImpl{
		TodoUsecase: tu,
	}
}

func (tc *TodoControllerImpl) FindById(ctx *gin.Context) {
	id, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todo := tc.TodoUsecase.FindById(ctx, id)

	ctx.JSON(http.StatusOK, todo)
}

func (tc *TodoControllerImpl) FindList(ctx *gin.Context) {
	todoList := tc.TodoUsecase.FindAll(ctx)
	ctx.JSON(http.StatusOK, todoList)
}

type CreateTodoDto struct {
	Title       string
	Description string
	UserId      uuid.UUID
}

func (tc *TodoControllerImpl) Create(ctx *gin.Context) {
	var body CreateTodoDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tc.TodoUsecase.Create(ctx, CreateTodoRequest(body))

	ctx.Status(http.StatusCreated)
}

type DeleteTodoDto struct {
	ID uuid.UUID
}

func (tc *TodoControllerImpl) Delete(ctx *gin.Context) {
	var body DeleteTodoDto
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tc.TodoUsecase.Delete(ctx, body.ID)
	ctx.JSON(http.StatusCreated, body.ID)
}
