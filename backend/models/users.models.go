package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID              int    `gorm:"primary_key" json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	Profile_picture string `json:"profile_picture"`
	Last_login      string `json:"last_login"`
	User_type       string `json:"user_type"`
}
