package services

import (
	"backend/db"
	"backend/models"
)

func NewUser(email string, name string) *models.User {
	return &models.User{Name: name, Email: email}
}

func AddUser(user *models.User) (string, error) {
	db.DB.Create(user)

	return user.ID, nil
}

func GetUserByID(id string) (*models.User, error) {
	user := &models.User{
		ID: id,
	}
	db.DB.
		//Preload("Completions.Challenge").
		Preload("Completions").
		Find(&user)

	for _, c := range user.Completions {
		c.Challenge = &models.Challenge{
			ID: c.ChallID,
		}
		db.DB.Find(&c.Challenge)
		user.Points += c.Challenge.Points
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{
		Email: email,
	}
	db.DB.
		//Preload("Completions.Challenge").
		Preload("Completions").
		Find(&user)

	for _, c := range user.Completions {
		c.Challenge = &models.Challenge{
			ID: c.ChallID,
		}
		db.DB.Find(&c.Challenge)
		user.Points += c.Challenge.Points
	}

	return user, nil
}
