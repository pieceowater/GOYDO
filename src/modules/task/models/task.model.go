package models

import "time"

// Task represents a task in the system
// swagger:model Task
type Task struct {
	// The unique identifier for the task
	// example: 1
	ID uint `json:"id"`

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
