package router

import (
	userrouter "test/router/user_router"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	userrouter.UserRouter(r)

	return r

}
