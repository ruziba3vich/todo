package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/todoGo/internal/authentication"
	"github.com/ruziba3vich/todoGo/internal/models"
	"github.com/ruziba3vich/todoGo/internal/services"
)

func CreateTask(c *gin.Context, db *sql.DB) {
	token := c.GetHeader("Authorization")

	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	id, err := authentication.ValidateJWTToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}
	var task models.Task
	c.ShouldBindJSON(&task)
	task.UserId = id
	createdTask, err := services.CreateTask(task, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Bad request"})
		return
	}
	c.JSON(http.StatusCreated, createdTask)
}
