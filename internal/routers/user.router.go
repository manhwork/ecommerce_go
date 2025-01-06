package routers

import (
	"github/manhwork/ecommerce_go/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	v1 := r.Group("/v1/api")
	{
		v1.GET("/users", controller.NewUserController().Index)
		v1.GET("/users/:cccd", controller.NewUserController().GetInfoUser)
		v1.GET("/", controller.NewUserController().Home)
		v1.GET("/search", controller.NewUserController().Search)
		v1.POST("/add", controller.NewUserController().Add)
	}
}
