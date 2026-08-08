package services

import (
	"errors"
	"strings"

	"github.com/FisumTeshome/Reciep_project/database"
	"github.com/FisumTeshome/Reciep_project/models"
	"github.com/FisumTeshome/Reciep_project/utils"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UpdateProfileInput struct {
	Username  string `json:"username"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Avatar    string `json:"avatar"`
	Bio       string `json:"bio"`
}

type ChangePasswordInput struct {
	CurrentPassword string `json:"current_password"`
	NewPassword     string `json:"new_password"`
}

func GetProfile(userID string) (models.User, error) {
	return GetUserByID(userID)
}

func UpdateProfile(userID string, input UpdateProfileInput) (models.User, error) {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", parsedUserID).Error; err != nil {
		return models.User{}, err
	}

	if strings.TrimSpace(input.Username) != "" && input.Username != user.Username {
		var existing models.User
		if err := database.DB.Where("username = ? AND id != ?", input.Username, parsedUserID).First(&existing).Error; err == nil {
			return models.User{}, errors.New("username is already taken")
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, err
		}
		user.Username = input.Username
	}

	user.FirstName = input.FirstName
	user.LastName = input.LastName
	user.Avatar = input.Avatar
	user.Bio = input.Bio

	if err := database.DB.Save(&user).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdatePassword(userID string, input ChangePasswordInput) error {
	parsedUserID, err := uuid.Parse(userID)
	if err != nil {
		return err
	}

	if strings.TrimSpace(input.NewPassword) == "" || len(input.NewPassword) < 6 {
		return errors.New("new password must be at least 6 characters")
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", parsedUserID).Error; err != nil {
		return err
	}

	if err := utils.ComparePassword(user.Password, input.CurrentPassword); err != nil {
		return errors.New("current password is incorrect")
	}

	hashedPassword, err := utils.HashPassword(input.NewPassword)
	if err != nil {
		return err
	}

	user.Password = hashedPassword
	return database.DB.Save(&user).Error
}