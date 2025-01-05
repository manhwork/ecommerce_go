package routers

import (
	"github/manhwork/ecommerce_go/internal/controller"

	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine) {
	v1 := r.Group("/v1/api")
	{
		v1.GET("/users", controller.NewUserController().Index)
	}
}
