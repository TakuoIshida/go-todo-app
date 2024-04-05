package todo

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ITodoUsecase interface {
	Create(ctx *gin.Context, req CreateTodoRequest)
	// Update(ctx *gin.Context, todo request.UpdatetodoRequest)
	Delete(ctx *gin.Context, id uuid.UUID)
	FindById(ctx *gin.Context, id uuid.UUID) Todo
	FindAll(ctx *gin.Context) []Todo
}

type CreateTodoRequest struct {
	Title       string
	Description string
	UserId      uuid.UUID
	// AttachmentFiles []AttachmentFile
}
