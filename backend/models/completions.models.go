package models

import (
	"backend/db"
)

type Completion struct {
	ChallID int    `gorm:"primary_key;column:chall_id" json:"chall_id"`
	UserID  string `gorm:"primary_key;column:user_id" json:"user_id"`

	Challenge *Challenge `gorm:"foreignkey:ID;references:ChallID" json:"challenge"`
}

func (c *Completion) Save() error {
	tx := db.DB.Save(c)
	return tx.Error
}
