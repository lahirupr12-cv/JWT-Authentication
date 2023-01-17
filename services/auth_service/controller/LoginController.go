package controller

import (
	"lahiruprasad12/services/auth_service/constants"
	"lahiruprasad12/services/auth_service/services"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context) string
}

type loginController struct {
	loginService services.LoginService
	jwtService   services.JWTService
}

func LoginHandler(loginService services.LoginService, jwtService services.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

func (controller *loginController) Login(ctx *gin.Context) string {
	var credetial constants.LoginCredentials
	err := ctx.ShouldBind(&credetial)

	if err != nil {
		return "not data"
	}

	isUserAuthenticated := controller.loginService.LogInUser(credetial.Email, credetial.Password)

	if isUserAuthenticated {
		return controller.jwtService.GenerateToken(credetial.Email, true)
	} else {
		return ""
	}
}
