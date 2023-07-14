package model

import "gorm.io/gorm"

type PaypayID struct {
	gorm.Model
	UserID   int `gorm:"foreignkey:ID"`
	PaypayID string
}
