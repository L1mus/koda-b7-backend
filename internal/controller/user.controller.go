package controller

import (
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/l1mus/koda-b7-backend/internal/dto"
)

type AuthController struct{}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (ac *AuthController) Login(ctx *gin.Context){
		var body dto.LoginRequest
		if err := ctx.ShouldBind(&body); err != nil {
			errStr := err.Error()

			if strings.Contains(errStr,"Email"){

			}

			log.Println("Error : ", err.Error())
			ctx.JSON(http.StatusInternalServerError,gin.H{
				"message" : "Internal server error",
			})
			return
		};
		validEmail := "limustadji@gmail.com"
		validPassword := "1234567"
		isValid, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`,body.Email)

		if !isValid {
			ctx.JSON(http.StatusBadRequest,gin.H{
				"message":"Email Invalid",
			})
			return
		}

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
	}

func (ac *AuthController) Register(ctx *gin.Context) {
		var body dto.RegisterRequest
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
	
		ctx.JSON(http.StatusOK,gin.H{
				"status" : "Success",
				"error" : "",
				"data" : dto.DataReponseRegister{
					Fullname: body.Fullname,
					Email: body.Email,
				},
				"message" : "Register berhasil !!!",
		})
	
}


