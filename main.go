package main

import (
	"github/manhwork/ecommerce_go/internal/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("templates/*")

	routers.NewRouter(r)

	r.Run(":8080")
}
