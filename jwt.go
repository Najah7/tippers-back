package main

import (
	"time"

	"github.com/golang-jwt/jwt"
)

func JwtGenerate() (string, error) {
	// Claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": 1, // TODO userIdを入れるようにする
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	tokenString, err := token.SignedString([]byte("SECRET_KEY"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
