package authentication

import (
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func ValidateJWTToken(tokenString string) (id string, e error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		return jwtSecret, nil
	})
	if err != nil {
		// fmt.Println(err)
		return id, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Println(claims, ok)
		userID, _ := claims["user_id"].(string)

		// if !ok {
		// 	fmt.Println("----------------------------------------")
		// 	fmt.Println(tokenString)
		// 	fmt.Println("----------------------------------------")
		// 	return id, errors.New("invalid token: user ID not found")
		// }
		if expired, err := CheckTokenExpiration(claims); expired || err != nil {
			fmt.Println("-------------------------------------------------------------------------came1")
			return userID, err
		}
		fmt.Println("AAAAAAAAAAAAAAAAAAAAAAAA", userID)
		return userID, nil
	}
	fmt.Println("-------------------------------------------------------------------------came3")
	return id, errors.New("invalid token")
}

func CheckTokenExpiration(claims jwt.MapClaims) (bool, error) {
	exp, ok := claims["exp"].(float64)
	// fmt.Println("checking ", exp, ok)
	if !ok {
		return false, errors.New("expiration time not found in token claims")
	}

	expirationTime := time.Unix(int64(exp), 0)
	if time.Now().After(expirationTime) {
		return true, nil
	}

	return false, nil
}
