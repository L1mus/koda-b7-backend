package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main () {
	route := gin.Default()
	route.POST("/register", func(ctx *gin.Context) {
		var body Register
		if err := ctx.ShouldBindWith(&body , binding.JSON); err != nil {
			log.Println("Error : ", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal server Error",
			})
		return
		}
		if body.Fullname == "" || body.Email == "" || body.Password == ""{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"message":"Input Tidak Boleh Kosong",
			})
			return
		}

		isMatch, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,body.Email)

		if !isMatch{
			ctx.JSON(http.StatusBadRequest,gin.H{
				"message":"Email Invalid",
			})
			return
		}

		type dataReponseRegister struct{
			Fullname, Email string
			
		}
		ctx.JSON(http.StatusOK,gin.H{
				"status" : "Success",
				"error" : "",
				"data" : dataReponseRegister{
					Fullname: body.Fullname,
					Email: body.Email,
				},
				"message" : "Register berhasil !!!",
		})
		var dataRegister []Register

		dataRegister = append(dataRegister, body)
		fmt.Println(dataRegister)

	})


	// Login Route
	route.POST("/login", func(ctx *gin.Context) {
		var body Login
		if err := ctx.ShouldBind(&body); err != nil {
			log.Println("Error : ", err.Error())
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"message" : "Internal server error",
			})
			return
		};
		validEmail := "limustadji@gmail.com"
		validPassword := "1234567"

		if body.Email == validEmail && body.Password == validPassword{
			ctx.JSON(http.StatusOK, gin.H{
				"status" : "Success",
				"error" : "",
				"data" : nil,
				"message" : "Login berhasil !!!",
			})
		}else{
			ctx.JSON(http.StatusOK, gin.H{
				"status" : "Success",
				"error" : "",
				"data" : nil,
				"message" : "Email atau Password Salah",
			})
		}
	})
	route.Run("localhost:8080")
}

type Login struct {
	Email string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

type Register struct {
	Fullname string `form:"fullname" json:"fullname"` 
	Email string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}