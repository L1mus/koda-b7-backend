package router

import (
	"github.com/gin-gonic/gin"
	"github.com/l1mus/koda-b7-backend/internal/controller"
)

func AuthRouter(route *gin.Engine) {
	authRouter := route.Group("/auth")
	authController := controller.NewAuthController()

	authRouter.POST("/", authController.Login)
	authRouter.POST("/register", authController.Register)
}