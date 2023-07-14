package model

import "gorm.io/gorm"

type Restaurant struct {
	gorm.Model
	Name            string
	ProfileImageURL string
	Description     string
}
