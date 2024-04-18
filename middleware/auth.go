package auth

import (
	"errors"
	"go-todo-app/feature/user"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AuthRequired() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Some authorization in Authorization
		var user user.UserContext
		// authorization logic
		user.TenantId, _ = uuid.Parse("6be98432-2812-4ce0-a342-214e52aa791c")
		user.Id, _ = uuid.Parse("a553b7ab-8af5-4c51-8e93-022c3d9f056a")
		user.AccountId, _ = uuid.Parse("a553b7ab-8af5-4c51-8e93-022c3d9f056a")
		user.Email = "example@gmail.com"
		user.LastName = "LastName"
		user.FirstName = "FirstName"
		ctx.Set("AuthorizedUser", user)
	}
}

// HandlerFunc
func GetPrincipal(ctx *gin.Context) (user.UserContext, error) {
	principal, exists := ctx.Get("AuthorizedUser")
	if !exists {
		return user.UserContext{}, errors.New("unauthorized")
	}
	userContext, ok := principal.(user.UserContext)
	if !ok {
		return user.UserContext{}, errors.New("invalid user type")
	}

	return userContext, nil
}
