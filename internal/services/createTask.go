package services

import (
	"database/sql"
	"time"

	"github.com/ruziba3vich/todoGo/internal/models"
)

func CreateTask(task models.Task, db *sql.DB) (models.Task, error) {
	task.CreatedOn = time.Now()
	query := "INSERT INTO tasks(user_id, name, title, content, createdOn, isCompleted) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;"
	var taskID int
	err := db.QueryRow(query, task.UserId, task.Name, task.Title, task.Content, task.CreatedOn, task.IsCompleted).Scan(&taskID)
	if err != nil {
		return models.Task{}, err
	}
	task.ID = taskID
	return task, nil
}
