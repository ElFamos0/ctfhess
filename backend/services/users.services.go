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
	var user models.User
	db.DB.
		Preload("Completions.Challenge").
		Preload("Completions").
		First(&user, id)

	for _, c := range user.Completions {
		user.Points += c.Challenge.Points
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	if err := db.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
