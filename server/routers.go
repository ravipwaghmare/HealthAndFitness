package server

import (
	"HealthAndFitness/routing"
)

var (
	UserRouter routing.User
)

//InitializeRouters will create routing page for user
func InitializeRouters(server Server) {
	UserRouter = routing.NewUserRouter(
		UserController,
	)
}
