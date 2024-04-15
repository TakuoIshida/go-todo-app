package todo

import (
	"go-todo-app/feature/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateTodoRequest struct {
	Title       string
	Description string
	// AttachmentFiles []AttachmentFile
}

type ITodoUsecase interface {
	Create(ctx *gin.Context, userContext user.UserContext, req CreateTodoRequest)
	// Update(ctx *gin.Context, todo request.UpdatetodoRequest)
	Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID)
	FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID) Todo
	FindAll(ctx *gin.Context, userContext user.UserContext) []Todo
}

type ITodoService interface {
	Create(ctx *gin.Context, userContext user.UserContext, t *Todo, session *gorm.DB)
	// Update(ctx *gin.Context, todo request.UpdatetodoRequest)
	Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB)
	FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) Todo
	FindAll(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []Todo
}

type ITodoRepository interface {
	Create(ctx *gin.Context, userContext user.UserContext, t *Todo, session *gorm.DB)
	// Update(ctx *gin.Context, todo Todo)
	Delete(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB)
	FindById(ctx *gin.Context, userContext user.UserContext, id uuid.UUID, session *gorm.DB) Todo
	FindAll(ctx *gin.Context, userContext user.UserContext, session *gorm.DB) []Todo
}
