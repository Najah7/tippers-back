package service

import (
	"os"
	"time"
	"tippers-back/db/table"

	"github.com/golang-jwt/jwt"
)

var jwtSecret = []byte(getJwtSecret())

func JwtGenerate(user table.User) (string, error) {
	// Claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func getJwtSecret() string {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return "secret"
	}
	return secret
}
