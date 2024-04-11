package main

import (
	"fmt"
	"go-todo-app/feature/todo"
	"go-todo-app/shared/database"
	"go-todo-app/shared/database/config"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	r := gin.New()
	if config.Conf.GoEnv == "release" {
		gin.SetMode(gin.ReleaseMode)
		fmt.Println("run in production")
	}
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	database.NewCommonClientConnector()
	tenantConn := database.NewTenantClientConnector()

	todoRepository := todo.NewTodoRepositoryImpl(tenantConn.DB)
	todoService := todo.NewTodoServiceImpl(todoRepository)
	todoUsecase := todo.NewTodoUsecaseImpl(todoService)
	todoController := todo.NewTodoController(todoUsecase)

	todoGroup := r.Group("/todo")
	{
		todoGroup.GET("/", todoController.FindList)
		todoGroup.GET("/:id", todoController.FindById)
		todoGroup.POST("/", todoController.Create)
		// todoGroup.PUT("/", todoController.Update)
		todoGroup.DELETE("/:id", todoController.Delete)
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Printf("Listen on http://localhost:%s !!!!!\n", config.Conf.Port)
	r.Run(fmt.Sprintf(":%s", config.Conf.Port))
}
