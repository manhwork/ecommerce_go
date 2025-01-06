package controller

import (
	"github/manhwork/ecommerce_go/internal/service"
	"net/http"

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
		c.HTML(http.StatusNotFound, "index.html", gin.H{
			"message": "Users not found",
		})
		return
	}

	c.HTML(http.StatusOK, "index.html", gin.H{
		"users": users,
	})
}

// [GET] /v1/api/users/:cccd
func (uc *UserController) GetInfoUser(c *gin.Context) {
	cccd := c.Param("cccd")
	user, err := uc.userService.FindUserById(cccd)
	if err != nil {
		c.HTML(404, "infouser.html", gin.H{
			"msg": "Error load info user",
		})
	}
	c.HTML(200, "infouser.html", gin.H{
		"user": user,
	})
}

// [GET] /v1/api/
func (uc *UserController) Home(c *gin.Context) {
	c.HTML(http.StatusOK, "home.html", nil)
}

// [GET] /v1/api/search
func (uc *UserController) Search(c *gin.Context) {

	key := c.Query("key")

	user, err := uc.userService.FindUserById(key)

	if err != nil {
		c.HTML(http.StatusNotFound, "search.html", gin.H{
			"user":   "Not Found User",
			"status": 404,
		})
		return
	}

	c.HTML(http.StatusOK, "search.html", gin.H{
		"user": user,
	})
}

// [GET] /v1/api/add
func (uc *UserController) Add(c *gin.Context) {

	c.JSON(201, gin.H{
		"msg": "Tao user",
	})
}
