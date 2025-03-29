package controllers

import (
	"test/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userSrv *services.UserService
}

func NewUserController(userSrv *services.UserService) *UserController {
	return &UserController{userSrv: userSrv}
}

func (uc *UserController) GetUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  uc.userSrv.GetUser(),
	})
}
