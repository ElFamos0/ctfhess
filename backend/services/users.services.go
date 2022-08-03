package services

import (
	"backend/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewUser(email string, name string) *models.User {
	return &models.User{Name: name, Email: email, User_type: "user"}
}

func AddUser(user *models.User) (int, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		print("Error opening database")
		return 0, err
	}

	db.Create(user)

	print("Im adding a user")
	return user.ID, nil
}

func GetUserByID(id int) (*models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var user models.User
	db.First(&user, id)

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	var user models.User

	if err := db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}
