package main

import (
	"github.com/gin-gonic/gin"
	"github.com/l1mus/koda-b7-backend/internal/router"
)

func main () {
	route := gin.Default()
	router.InitRoute(route)
	route.Run("localhost:8080")
}
