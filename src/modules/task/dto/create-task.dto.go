package dto

import "time"

// CreateTaskDTO represents the data required to create a new task.
// swagger:model CreateTaskDTO
type CreateTaskDTO struct {
	// The title of the task
	// example: "Task Title"
	Title string `json:"title" binding:"required"`

	// The description of the task
	// example: "This is a detailed description of the task."
	Description *string `json:"description"`

	// The deadline for the task
	// example: "2024-12-31"
	Deadline *time.Time `json:"deadline"`

	// Indicates if the task is urgent
	// example: true
	Urgent *bool `json:"urgent"`
}
