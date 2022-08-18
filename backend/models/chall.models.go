package models

import "gorm.io/gorm"

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
	gorm.Model
	ID               int    `gorm:"primary_key" json:"id"`
	Title            string `json:"title"`
	Subtitle         string `json:"subtitle"`
	ShortDescription string `json:"short_description"`
	OpensAt          string `json:"day_open"`
	Flag             string `json:"flag"`
	Points           int    `json:"points"`

	Pages []*ChallengePage `json:"pages"` // Il faut preload avec le bon ordre (ID asc)
}

type ChallengePage struct {
	gorm.Model
	ID          int    `gorm:"primary_key" json:"id"`
	ChallengeID int    `json:"challenge_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	AskFlag     bool   `json:"flag"`
}
