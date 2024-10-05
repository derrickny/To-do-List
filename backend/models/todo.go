package models

// Status represents the state of a Todo item. It can be either:
//   - "pending": the task has not been started yet
//   - "completed": the task is finished
//   - "ongoing": the task is currently being worked on
type Status string

// These constants represent the possible states of a Todo item
const (
	StatusPending   Status = "pending"   // The task has not been started yet
	StatusCompleted Status = "completed" // The task is finished
	StatusOngoing   Status = "ongoing"   // The task is currently being worked on
)

// Todo represents a single Todo item
type Todo struct {
	// ID is a unique identifier for the Todo item
	ID uint `json:"id" gorm:"primary_key"`

	// Title is the title of the Todo item
	Title string `json:"title"`

	// Status is the current state of the Todo item
	// The default value is "pending", meaning the task has not been started yet
	Status Status `json:"status" gorm:"default:'pending'"`
}
