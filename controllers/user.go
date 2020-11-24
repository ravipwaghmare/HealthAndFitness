package controllers

import (
	"HealthAndFitness/model"
	"strings"

	"gorm.io/gorm"
)

// User controller
type User struct {
	database *gorm.DB
}

// NewUserController will create new user controller
func NewUserController(database *gorm.DB) User {
	return User{
		database: database,
	}
}

// Create will make a new user
func (controller User) Create(user *model.User) error {
	user.Email = strings.ToLower(user.Email)
	return controller.database.Create(user).Error
}
