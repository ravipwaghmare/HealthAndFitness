package routing

import "HealthAndFitness/interfaces"

// User struct
type User struct {
	userController interfaces.User
}

// NewSignInRouter function
func NewUserRouter(
	userController interfaces.User,
) User {
	return User{
		userController: userController,
	}
}
