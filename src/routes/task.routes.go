package routes

import (
	"GOYDO/src/modules/task"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(r *gin.Engine, db *gorm.DB) {
	taskService := &task.Service{DB: db}
	taskController := &task.Controller{Service: *taskService}

	taskRoutes := r.Group("/tasks")
	{
		taskRoutes.GET("/", taskController.GetTasks)
		taskRoutes.POST("/", taskController.CreateTask)
		taskRoutes.GET("/:id", taskController.GetTask)
		taskRoutes.PUT("/:id", taskController.UpdateTask)
		taskRoutes.DELETE("/:id", taskController.DeleteTask)
	}
}
