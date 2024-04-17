package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID          int       `json:"id"`
	UserId      uuid.UUID `json:"user_id"`
	Name        string    `json:"name"`
	Title       string    `json:"title"`
	Content     string    `json:"content"`
	CreatedOn   time.Time `json:"created_on"`
	CompletedOn time.Time `json:"completed_on"`
	IsCompleted bool      `json:"is_completed"`
}
