package interfaces

import (
	"HealthAndFitness/model"
)

// User interface
type User interface {
	Create(user *model.User) error
	FindByEmail(email string) (*model.User, error)
}
