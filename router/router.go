package router

import (
	di_wire "test/di"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	userController := di_wire.InitUserRouterHandler()

	v1 := r.Group("/v1")
	{
		v1.GET("/user", userController.GetUser) // 🛠 Dependency Injection
	}

	return r

}
