package authentication

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte("f350124f5d1497b709735c3cc842041b3534c0e1a87f0ba0abd5c37ee8369bf4")

func GenerateJWTToken(userID string) (string, error) {
	expirationDays := 60
	expirationTime := time.Now().Add(time.Duration(24*expirationDays) * time.Hour)

	claims := &jwt.StandardClaims{
		ExpiresAt: expirationTime.Unix(),
		Subject:   userID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
