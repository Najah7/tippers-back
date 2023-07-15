package test

import (
	"testing"
	"tippers-back/service"
)

func TestPassword(t *testing.T) {
	password := "password1234"
	hashPassword := service.HashPassword(password)
	if service.ComparePassword(hashPassword, password) != nil {
		t.Error("Invalid password")
		return
	}

}
