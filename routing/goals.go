package routing

import (
	"HealthAndFitness/interfaces"
	"net/http"

	"github.com/labstack/echo"
)

// Goals struct
type Goals struct {
	goalsController interfaces.Goals
}

// NewGoalsRouter function
func NewGoalsRouter(
	goalsController interfaces.Goals,
) Goals {
	return Goals{
		goalsController: goalsController,
	}
}

// Register new goals router
func (router Goals) Register(group *echo.Group) {
	group.PUT("/weight", router.setweight)
	group.PUT("/steps", router.setsteps)
}

func (router Goals) setweight(context echo.Context) error {
	request := &struct {
		ID     int `json:"id" validator:"required"`
		Weight int `json:"weight" validator:"required"`
	}{}

	// bind context to request
	if err := context.Bind(request); err != nil {
		return err
	}

	// Add weight in the database
	err := router.goalsController.SetWeightGoal(request.ID, request.Weight)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{
		"message": "Weight Goals Added Successfully",
	})
}

func (router Goals) setsteps(context echo.Context) error {
	request := &struct {
		ID    int `json:"id" validator:"required"`
		Steps int `json:"steps" validator:"required"`
	}{}

	// bind context to request
	if err := context.Bind(request); err != nil {
		return err
	}

	// Add weight in the database
	err := router.goalsController.SetStepGoals(request.ID, request.Steps)
	if err != nil {
		return err
	}

	return context.JSON(http.StatusOK, map[string]string{
		"message": "Steps Goals Added Successfully",
	})
}
