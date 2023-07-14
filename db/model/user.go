package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email            string `gorm:"unique"`
	Name             string
	Password         string
	ProfileImageURL  string
	Dream            string
	Major            string
	IsStudent        int
	IsEmployed       int
	PeriodOfWorkings string
	RestaurantID     int `gorm:"foreignkey:ID"`
}
