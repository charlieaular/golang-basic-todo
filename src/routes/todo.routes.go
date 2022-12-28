package routes

import (
	"golang-todo/src/handlers"

	"golang-todo/src/repositories"
	"golang-todo/src/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisteTodoRoutes(router *gin.Engine, db *gorm.DB) {

	todoRepo := repositories.NewTodoRepo(db)
	todoService := services.TodoService(todoRepo)
	todoHandler := handlers.NewTodoHandler(todoService)

	todoRoutes := router.Group("todo")
	{
		todoRoutes.GET("/", todoHandler.GetAll)
		todoRoutes.POST("/", todoHandler.Create)
		todoRoutes.PUT("/change-visibility/:id", todoHandler.ToggleDone)
		todoRoutes.PUT("/change-all-done", todoHandler.ChangeAllDone)
		todoRoutes.GET("/:id", todoHandler.Get)
		todoRoutes.PUT("/:id", todoHandler.Update)
		todoRoutes.DELETE("/:id", todoHandler.Delete)
	}

}
