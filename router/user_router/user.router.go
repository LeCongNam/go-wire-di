package userrouter

import (
	di_wire "test/di"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.Engine) *gin.RouterGroup {
	userController := di_wire.InitUserRouterHandler()

	userRouter := r.Group("/user")
	{
		userRouter.GET("/", userController.GetUser)
	}

	return userRouter
}
