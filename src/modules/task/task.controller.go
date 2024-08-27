package task

import (
	"GOYDO/src/modules/task/dto"
	"GOYDO/src/modules/task/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

// TaskController handles operations related to tasks.
type TaskController struct {
	Service TaskService
}

// CreateTask godoc
// @Summary      Create a task
// @Description  Create a new task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        task body dto.CreateTaskDTO true "Task"
// @Success      201 {object} models.Task
// @Failure      400 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /tasks [post]
func (ctrl *TaskController) CreateTask(c *gin.Context) {
	var createTaskDTO dto.CreateTaskDTO
	if err := c.ShouldBindJSON(&createTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		Title:       createTaskDTO.Title,
		Description: createTaskDTO.Description,
		Deadline:    createTaskDTO.Deadline,
		Urgent:      createTaskDTO.Urgent,
	}

	createdTask, err := ctrl.Service.CreateTask(task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}

// GetTasks godoc
// @Summary      Get all tasks
// @Description  Retrieve all tasks from the database
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Success      200 {array} models.Task
// @Failure      500 {object} map[string]string
// @Router       /tasks [get]
func (ctrl *TaskController) GetTasks(c *gin.Context) {
	tasks, err := ctrl.Service.GetTasks()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, tasks)
}

// GetTask godoc
// @Summary      Get a task by ID
// @Description  Retrieve a single task by its ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id path string true "Task ID"
// @Success      200 {object} models.Task
// @Failure      404 {object} map[string]string
// @Router       /tasks/{id} [get]
func (ctrl *TaskController) GetTask(c *gin.Context) {
	id := c.Param("id")

	task, err := ctrl.Service.GetTask(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// UpdateTask godoc
// @Summary      Update a task
// @Description  Update an existing task
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id path string true "Task ID"
// @Param        task body dto.UpdateTaskDTO true "Task"
// @Success      200 {object} models.Task
// @Failure      400 {object} map[string]string
// @Failure      404 {object} map[string]string
// @Failure      500 {object} map[string]string
// @Router       /tasks/{id} [put]
func (ctrl *TaskController) UpdateTask(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid task ID"})
		return
	}

	var updateTaskDTO dto.UpdateTaskDTO
	if err := c.ShouldBindJSON(&updateTaskDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	task := models.Task{
		ID:          uint(id),
		Title:       *updateTaskDTO.Title,
		Description: updateTaskDTO.Description,
		Deadline:    updateTaskDTO.Deadline,
		Urgent:      updateTaskDTO.Urgent,
	}

	updatedTask, err := ctrl.Service.UpdateTask(idStr, task)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedTask)
}

// DeleteTask godoc
// @Summary      Delete a task
// @Description  Delete a task by its ID
// @Tags         tasks
// @Accept       json
// @Produce      json
// @Param        id path string true "Task ID"
// @Success      204
// @Failure      404 {object} map[string]string
// @Router       /tasks/{id} [delete]
func (ctrl *TaskController) DeleteTask(c *gin.Context) {
	id := c.Param("id")

	if err := ctrl.Service.DeleteTask(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
