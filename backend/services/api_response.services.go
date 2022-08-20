package services

import (
	"backend/models"
)

func NewProfileDataResponse(user *models.User) *models.ProfileDataResponse {
	return &models.ProfileDataResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
