package routing

import (
	"HealthAndFitness/interfaces"
	"HealthAndFitness/model"
	"fmt"
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
	group.POST("/signup", router.signUp)
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

func (router User) signUp(context echo.Context) error {
	request := &struct {
		UserName string `json:"username" validator:"required"`
		Name     string `json:"name" validator:"required"`
		Email    string `json:"email" validator:"required,email"`
		Password string `json:"password" validator:"required"`
		Gender   string `json:"gender" validator:"required"`
		DOB      string `json:"dob" validator:"required"`
		Weight   int    `json:"weight" validator:"required"`
		Height   int    `json:"height" validator:"required"`
	}{}

	// bind context to request
	if err := context.Bind(request); err != nil {
		return err
	}

	var bmi float32
	heightInMeter := float32(request.Height) / 100
	weightinfloat := float32(request.Weight)
	fmt.Println("Weight in float", weightinfloat)
	fmt.Println("Height in float", heightInMeter)
	bmi = (weightinfloat / (heightInMeter * heightInMeter))

	user := &model.User{
		UserName:      request.UserName,
		Name:          request.Name,
		Email:         request.Email,
		Password:      request.Password,
		Gender:        request.Gender,
		DOB:           request.DOB,
		Weight:        request.Weight,
		Height:        request.Height,
		AccountStatus: 1,
		BMI:           bmi,
	}

	// Add user in the database
	err := router.userController.Create(user)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{
		"message": "User Added Successfully",
	})

}
