package router

import (
	userHandler "gin-api-sample/internal/user/interface/handler"

	"github.com/gin-gonic/gin"
	// taskHandler "gin-api-sample/internal/task/interface/handler"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Userハンドラーの設定
	userHandler := userHandler.NewUserHandler()
	r.GET("/users", userHandler.GetUsers)
	r.POST("/users", userHandler.CreateUser)
	r.PUT("/users/:id", userHandler.UpdateUser)
	r.DELETE("/users/:id", userHandler.DeleteUser)

	// // Taskハンドラーの設定
	// taskHandler := taskHandler.NewTaskHandler()
	// r.GET("/tasks", taskHandler.GetTasks)
	// r.POST("/tasks", taskHandler.CreateTask)
	// r.PUT("/tasks/:id", taskHandler.UpdateTask)
	// r.DELETE("/tasks/:id", taskHandler.DeleteTask)

	return r
}
