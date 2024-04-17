package services

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/ruziba3vich/todoGo/internal/authentication"
	"github.com/ruziba3vich/todoGo/internal/models"
)

func Register(user models.User, db *sql.DB) (string, error) {
	user.ID = uuid.NewString()
	query := "INSERT INTO users(id, username, password) VALUES ($1, $2, $3);"
	var erro error
	user.Password, erro = authentication.HashPassword(user.Password)
	if erro != nil {
		return "", erro
	}
	if _, err := db.Exec(query, user.ID, user.Username, user.Password); err != nil {
		return "", err
	}
	token, err := authentication.GenerateJWTToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
