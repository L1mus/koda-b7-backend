package middleware

import (
	"net/http"
	"slices"
	"strings"

	"github.com/gin-gonic/gin"
)

func CORSMiddleware(ctx *gin.Context){
	allowedOrigin := []string{"http://127.0.0.1:5500"}
	currentOrigin := ctx.GetHeader("Origin")
	if slices.Contains(allowedOrigin,currentOrigin){
		ctx.Header("Access-Control-Allow-Origin",currentOrigin)
	}

	allowedHeader := []string{
		"Content-Type","X-koda-X",
	}
	ctx.Header("Access-Control-Allow-Headers",strings.Join(allowedHeader,", "))
	
	allowedMethod := []string{http.MethodGet,http.MethodPost,http.MethodPut,http.MethodPatch,http.MethodDelete,}

	ctx.Header("Accces-Control-Allow-Methods",strings.Join(allowedMethod,", "))

	if ctx.Request.Method == http.MethodOptions{
		ctx.AbortWithStatus(http.StatusNoContent)
		return
	}

	ctx.Next()
}