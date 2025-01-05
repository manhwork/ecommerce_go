package controller

import (
	"github/manhwork/ecommerce_go/internal/service"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController() *UserController {
	return &UserController{
		userService: service.NewUserService(),
	}
}

// /v1/api/users
func (uc *UserController) Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"msg": "ok",
	})
}
