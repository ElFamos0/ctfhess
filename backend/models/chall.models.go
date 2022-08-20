package models

import (
	"backend/db"
)

// title: '',
// subtitle: '',
// short_description: '',
// day_open: '',
// flag: '',
// points: '',
// pages: [
//   {
// 	title: '',
// 	description: '',
// 	flag: false,
//   },
// ],

type Challenge struct {
	ID               int    `gorm:"primary_key" json:"id"`
	Title            string `json:"title"`
	Subtitle         string `json:"subtitle"`
	ShortDescription string `json:"short_description"`
	Flag             string `json:"flag"`
	OpensAt          int    `json:"day_open"`
	Points           int    `json:"points"`
	Fake             bool   `gorm:"-" json:"fake"`

	Pages []*ChallengePage `json:"pages"` // Il faut preload avec le bon ordre (ID asc)
}

type ChallengePage struct {
	ID          int    `gorm:"primary_key" json:"id"`
	ChallengeID int    `gorm:"primary_key" json:"challenge_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AskFlag     bool   `json:"flag"`
}

func (c *Challenge) Save() error {
	for i, p := range c.Pages {
		p.ID = i
	}
	tx := db.DB.Save(c)
	return tx.Error
}

func (p *ChallengePage) Save() error {
	tx := db.DB.Save(p)
	return tx.Error
}

func (p *ChallengePage) Delete() error {
	tx := db.DB.Delete(p)
	return tx.Error
}

func (c *Challenge) Delete() error {
	db.DB.Table("completions").Delete(&Completion{}, "chall_id = ?", c.ID)
	db.DB.Table("challenge_pages").Delete(&ChallengePage{}, "challenge_id = ?", c.ID)

	tx := db.DB.Delete(c)
	return tx.Error
}

func GetAllChallenges() ([]*Challenge, error) {
	var challenges []*Challenge
	err := db.DB.
		Preload("Pages").
		//Where("id <= ?", jourJ)
		Find(&challenges).Error
	return challenges, err
}

func GetChallenge(ID int) (*Challenge, error) {
	challenge := &Challenge{
		ID: ID,
	}
	err := db.DB.
		Preload("Pages").
		//Where("id <= ?", jourJ)
		Find(&challenge).Error
	return challenge, err
}
