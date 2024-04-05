package main

import (
	"go-todo-app/feature/todo"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	// conn := database.NewDBClientConnector()

	todoRepository := todo.NewTodoRepositoryImpl(conn.DB)
	todoService := todoserviceimpl.NewTodoServiceImpl(todoRepository)
	todoUsecase := todousecaseimpl.NewTodoUsecaseImpl(todoService)
	todoController := controller.NewTodoController(todoUsecase)

	todoGroup := router.Group("/todo")
	{
		todoGroup.GET("/", todoController.FindList)
		todoGroup.GET("/:id", todoController.FindById)
		todoGroup.POST("/", todoController.Create)
		// todoGroup.PUT("/", todoController.Update)
		todoGroup.DELETE("/", todoController.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run()
}
