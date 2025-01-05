package main

import (
	"github/manhwork/ecommerce_go/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	routers.NewRouter(r)

	r.Run(":8080")
}
