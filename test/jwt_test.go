package test

import (
	"encoding/base64"
	"encoding/json"
	"strings"
	"testing"
	"tippers-back/db/model"
	"tippers-back/service"
)

func TestJwtGenerate(t *testing.T) {
	type Claims struct {
		UserID int `json:"user_id"`
	}

	user := model.User{}
	user.ID = 1
	token, err := service.JwtGenerate(user)
	if err != nil {
		t.Errorf("toke作成失敗 token: %v", token)
		return
	}
	validate, err := service.JwtValidate(token)
	if err != nil || !validate.Valid {
		t.Errorf("token 検証失敗 token: %v", token)
		return
	}
	// トークンをドットで分割
	tokenParts := strings.Split(token, ".")
	if len(tokenParts) != 3 {
		t.Error("Invalid token format")
		return
	}

	// ペイロード部分をデコード
	payload, err := base64.RawURLEncoding.DecodeString(tokenParts[1])
	if err != nil {
		t.Error("Failed to decode payload:", err)
		return
	}

	claims := struct {
		UserID int
	}{
		UserID: int(user.ID),
	}
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		t.Error("Failed to parse payload:", err)
		return
	}

	if int(user.ID) != claims.UserID {
		t.Errorf("want: %v got: %v", int(user.ID), claims.UserID)
	}
}
