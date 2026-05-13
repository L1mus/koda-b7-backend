package router

import "github.com/gin-gonic/gin"

func InitRoute(route *gin.Engine) {
	AuthRouter(route)
}