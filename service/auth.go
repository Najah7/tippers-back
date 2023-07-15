package service

import (
	"fmt"
	"tippers-back/db"
	"tippers-back/db/model"

	"github.com/golang-jwt/jwt"
)

// Retrieve the user from the database
func GetUserByName(db db.DB, name string) (*model.User, error) {
	user, err := db.GetUserByMail(name)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// Check if the password is correct
func VerifyPassword(userPassword, inputPassword string) bool {
	if ComparePassword(userPassword, inputPassword) != nil {
		return false
	}

	return true
}

func JwtValidate(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there's a problem with the signing method")
		}
		return jwtSecret, nil
	})
}
