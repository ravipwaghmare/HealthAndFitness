package controllers

import (
	"HealthAndFitness/model"
	"errors"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

// Goals controller
type Goals struct {
	database *gorm.DB
}

// NewGoalsController will create new goals controller
func NewGoalsController(database *gorm.DB) Goals {
	return Goals{
		database: database,
	}
}

// SetWeightGoal will set goals for weight
func (controller Goals) SetWeightGoal(id int, weight int) error {

	if err := controller.database.Model(&model.User{}).Where("id = ?", id).Update("weight_goal", weight).Error; err != nil {
		log.Error(err)
		return errors.New("Error while updating weight goal")
	}

	return nil
}

// SetStepGoals will set goals for steps
func (controller Goals) SetStepGoals(id int, steps int) error {

	if err := controller.database.Model(&model.User{}).Where("id = ?", id).Update("steps_goal", steps).Error; err != nil {
		log.Error(err)
		return errors.New("Error while updating steps goal")
	}

	return nil
}
