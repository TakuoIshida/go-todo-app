package todo

import (
	auth "go-todo-app/middleware"
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

func (tc *TodoControllerImpl) Create(ctx *gin.Context) {
	var body struct {
		Title       string
		Description string
		UserId      uuid.UUID
		TenantId    uuid.UUID
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tc.TodoUsecase.Create(ctx, CreateTodoRequest(body))

	ctx.Status(http.StatusCreated)
}

func (tc *TodoControllerImpl) Delete(ctx *gin.Context) {
	var body struct {
		Id string
	}
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validId, err := uuid.Parse(body.Id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userContext, err := auth.GetPrincipal(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tc.TodoUsecase.Delete(ctx, userContext, validId)
	ctx.JSON(http.StatusNoContent, validId)
}
