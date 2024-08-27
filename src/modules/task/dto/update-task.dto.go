package dto

import "time"

// UpdateTaskDTO represents the data required to update an existing task.
// swagger:model UpdateTaskDTO
type UpdateTaskDTO struct {
	// The title of the task
	// example: "Updated Task Title"
	Title *string `json:"title"`

	// The description of the task
	// example: "This is an updated description of the task."
	Description *string `json:"description"`

	// The deadline for the task
	// example: "2024-12-31"
	Deadline *time.Time `json:"deadline"`

	// Indicates if the task is urgent
	// example: true
	Urgent *bool `json:"urgent"`
}
