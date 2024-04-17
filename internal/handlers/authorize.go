package handlers

import (
	"database/sql"
	"log"
	"net/http"

	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/todoGo/internal/authentication"
)

func Authorize(c *gin.Context, db *sql.DB) {
	token := c.GetHeader("Authorization")
	if len(token) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "user unauthorized"})
		return
	}

	const bearerPrefix = "Bearer "
	if !strings.HasPrefix(token, bearerPrefix) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token format"})
		return
	}

	token = token[len(bearerPrefix):]

	id, err := authentication.ValidateJWTToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
		return
	}

	newToken, err := authentication.GenerateJWTToken(id)
	if err != nil {
		log.Fatal(err)
		return
	}

	c.JSON(http.StatusAccepted, newToken)
}
