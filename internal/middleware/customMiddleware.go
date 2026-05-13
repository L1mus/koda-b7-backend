package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CustomeMyCustomMiddleware(ctx *gin.Context) {
	xkodax := ctx.GetHeader("X-Koda-X")
	if xkodax != "aku koda"{
		ctx.AbortWithStatus(http.StatusConflict)
	}
	ctx.Next()
}