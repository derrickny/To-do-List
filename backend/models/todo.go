package models

type Todo struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Title  string `json:"title"`
	Status string `json:"status"` // pending, in_progress, completed, cancelled
}
