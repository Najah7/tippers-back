package table

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email            string `gorm:"unique" json:"email"`
	Name             string `json:"name"`
	Password         string `json:"password"`
	ProfileImageURL  string `json:"profileImageURL"`
	Dream            string `json:"dream"`
	Major            string `json:"major"`
	Money            int    `json:"money"`
	IsStudent        int    `json:"isStudent"`
	IsEmployed       int    `json:"isEmployed"`
	PeriodOfWorkings string `json:"periodOfWorkings"`
	RestaurantID     int    `gorm:"foreignkey:ID" json:"restaurantID"`
}
