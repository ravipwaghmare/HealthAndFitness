package routing

import (
	"HealthAndFitness/interfaces"
	"HealthAndFitness/model"
	"net/http"
	"strings"

	"github.com/labstack/echo"
)

// User struct
type User struct {
	userController interfaces.User
}

// NewUserRouter function
func NewUserRouter(
	userController interfaces.User,
) User {
	return User{
		userController: userController,
	}
}

// Register new signin router
func (router User) Register(group *echo.Group) {
	group.POST("/signin", router.signIn)
}

func (router User) signIn(context echo.Context) error {
	request := &struct {
		Email    string `json:"email" validator:"required,email"`
		Password string `json:"password" validator:"required,email"`
	}{}

	// bind context to request
	if err := context.Bind(request); err != nil {
		return err
	}

	// find user in database by email
	user, err := router.userController.FindByEmail(strings.TrimSpace(request.Email))
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, model.SignInResponse{
		User: *user,
	})

}
