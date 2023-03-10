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
	ID             string   `gorm:"primary_key" json:"id"`
	Email          string   `gorm:"column:email" json:"email"`
	Name           string   `gorm:"column:name" json:"name"`
	Surname        string   `gorm:"column:surname" json:"surname"`
	Promo          int      `gorm:"column:promo" json:"promo"`
	Spe            string   `gorm:"column:spe" json:"spe"`
	Type           UserType `gorm:"column:type" json:"type"`
	LastSubmission int64    `gorm:"column:last_submission" json:"last_submission"`

	Completions []*Completion `gorm:"foreignkey:UserID" json:"completions"`
	Points      int           `gorm:"-" json:"points"`

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

func (u *User) GetPoints() (pt int) {
	for _, c := range u.Completions {
		c.Challenge = &Challenge{
			ID: c.ChallID,
		}
		db.DB.Find(&c.Challenge)
		pt += c.Challenge.Points
	}

	return pt
}

func GetAdminList() (admins []*User) {
	db.DB.
		Preload("Completions").
		Find(&admins, "type = ?", 1)
	return
}

func GetUserList() (users []*User) {
	db.DB.
		Preload("Completions").
		Find(&users, "type = ?", 0)
	return
}

func (u *User) MakeAdmin() {
	u.Type = AdminUser
	u.Save()
}

func (u *User) MakeNormal() {
	u.Type = NormalUser
	u.Save()
}

func GetUserByID(id string) (user *User) {
	db.DB.First(&user, "id = ?", id)
	return
}
