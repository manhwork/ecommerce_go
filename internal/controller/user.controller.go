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

// [GET] /v1/api/users
func (uc *UserController) Index(c *gin.Context) {
	users, err := uc.userService.FindAllUsers()
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "NOT FOUND USERS!",
		})
	}
	c.JSON(200, users)
}

// [GET] /v1/api/users/:cccd
func (uc *UserController) GetInfoUser(c *gin.Context) {
	cccd := c.Param("cccd")
	user, err := uc.userService.FindUserById(cccd)
	if err != nil {
		c.JSON(404, gin.H{
			"msg": "NOT FOUND USERS!",
		})
	}
	c.JSON(200, user)
}
