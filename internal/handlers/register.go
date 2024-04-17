package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/todoGo/internal/models"
	"github.com/ruziba3vich/todoGo/internal/services"
)

func Register(c *gin.Context, db *sql.DB) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "invalid user information"})
		return
	}
	token, err := services.Register(user, db)
	if err != nil {
		c.JSON(http.StatusBadGateway, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"token":token})
}
