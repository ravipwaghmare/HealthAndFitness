package server

import (
	"HealthAndFitness/routing"
)

var (
	UserRouter routing.User
	GoalRouter routing.Goals
)

//InitializeRouters will create routing page for user
func InitializeRouters(server Server) {
	UserRouter = routing.NewUserRouter(
		UserController,
	)

	// register routes
	authGroup := server.Group("/auth")
	UserRouter.Register(authGroup)
	UserRouter.Register(authGroup.Group("/user"))

	GoalRouter = routing.NewGoalsRouter(
		GoalsController,
	)

	GoalRouter.Register(server.Group("/goals"))

}
