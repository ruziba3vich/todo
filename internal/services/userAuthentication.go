package services

import (
	"database/sql"
	"errors"

	"github.com/ruziba3vich/todoGo/internal/authentication"
	"github.com/ruziba3vich/todoGo/internal/models"
)

func Authenticate(username, password string, db *sql.DB) (string, error) {
	var user models.User
	query := "SELECT id, password FROM users WHERE username = $1"
	err := db.QueryRow(query, username).Scan(&user.ID, &user.Password)
	if err != nil {
		return "", errors.New("user is not found")
	}

	if !authentication.CheckPassword(password, user.Password) {
		return "", errors.New("password did not match")
	}

	token, err := authentication.GenerateJWTToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
