package models

import (
	"backend/db"
)

const (
	NormalUser UserType = iota
	AdminUser
)

type UserType int

type User struct {
	ID      string   `gorm:"primary_key" json:"id"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Surname string   `json:"surname"`
	Promo   int      `json:"promo"`
	Spé     string   `json:"spé"`
	Type    UserType `json:"type"`

	Completions []*Completion `gorm:"foreignkey:UserID" json:"completions"`

	// Name            string `json:"name"`
	// Email           string `json:"email"`
	// Profile_picture string `json:"profile_picture"`
	// Last_login      string `json:"last_login"`
	// User_type       string `json:"user_type"`
}

func (u *User) Save() error {
	tx := db.DB.Save(u)
	return tx.Error
}

func (u *User) Profile() *ProfileDataResponse {
	return &ProfileDataResponse{
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
	}
}
