package server

import (
	"HealthAndFitness/controllers"

	"gorm.io/gorm"
)

var (
	UserController  controllers.User
	GoalsController controllers.Goals
)

//InitializeControllers initialize the controllers to communicate
func InitializeControllers(db *gorm.DB) {
	UserController = controllers.NewUserController(db)
	GoalsController = controllers.NewGoalsController(db)
}
