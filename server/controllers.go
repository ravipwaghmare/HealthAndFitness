package server

import (
	"HealthAndFitness/controllers"

	"gorm.io/gorm"
)

var (
	UserController controllers.User
)

//InitializeControllers initialize the controllers to communicate
func InitializeControllers(db *gorm.DB) {
	UserController = controllers.NewUserController(db)
}
