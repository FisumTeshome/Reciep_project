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

type RegisterInput struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Avatar     string `json:"avatar"`
	Bio        string `json:"bio"`
}

type LoginInput struct {
	Identifier string `json:"identifier"`
	Password   string `json:"password"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  models.User `json:"user"`
}

func RegisterUser(input RegisterInput) (models.User, string, error) {
	if strings.TrimSpace(input.Username) == "" || strings.TrimSpace(input.Password) == "" || strings.TrimSpace(input.Email) == "" {
		return models.User{}, "", errors.New("username, email and password are required")
	}

	var existing models.User
	if err := database.DB.Where("username = ?", input.Username).First(&existing).Error; err == nil {
		return models.User{}, "", errors.New("username is already taken")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, "", err
	}

	if err := database.DB.Where("email = ?", input.Email).First(&existing).Error; err == nil {
		return models.User{}, "", errors.New("email is already registered")
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		return models.User{}, "", err
	}

	hashedPassword, err := utils.HashPassword(input.Password)
	if err != nil {
		return models.User{}, "", err
	}

	user := models.User{
		Username:  input.Username,
		Email:     input.Email,
		Password:  hashedPassword,
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Avatar:    input.Avatar,
		Bio:       input.Bio,
	}

	if err := database.DB.Create(&user).Error; err != nil {
		return models.User{}, "", err
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil
}

func LoginUser(input LoginInput) (models.User, string, error) {
	if strings.TrimSpace(input.Identifier) == "" || strings.TrimSpace(input.Password) == "" {
		return models.User{}, "", errors.New("identifier and password are required")
	}

	var user models.User
	if err := database.DB.Where("email = ? OR username = ?", input.Identifier, input.Identifier).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return models.User{}, "", errors.New("invalid credentials")
		}
		return models.User{}, "", err
	}

	if err := utils.ComparePassword(user.Password, input.Password); err != nil {
		return models.User{}, "", errors.New("invalid credentials")
	}

	token, err := utils.GenerateToken(user)
	if err != nil {
		return models.User{}, "", err
	}

	return user, token, nil
}

func GetUserByID(id string) (models.User, error) {
	parsed, err := uuid.Parse(id)
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	if err := database.DB.First(&user, "id = ?", parsed).Error; err != nil {
		return models.User{}, err
	}
	return user, nil
}

func GetUserFromToken(tokenString string) (models.User, error) {
	claims, err := utils.ValidateToken(tokenString)
	if err != nil {
		return models.User{}, err
	}
	return GetUserByID(claims.UserID)
}

func ExtractBearerToken(authHeader string) (string, error) {
	if strings.TrimSpace(authHeader) == "" {
		return "", errors.New("missing authorization header")
	}
	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
		return "", errors.New("invalid authorization header format")
	}
	return parts[1], nil
}