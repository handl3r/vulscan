package controllers

import (
	"github.com/gin-gonic/gin"
	"vulscan/api/http/context"
	"vulscan/src/packages"
)

type AuthenticationController struct {
	baseController
}

func NewAuthenticationController(appContext *context.ApplicationContext) *AuthenticationController {
	return &AuthenticationController{
		baseController{
			AppContext: appContext,
		},
	}
}

func (a *AuthenticationController) Register(c *gin.Context) {
	var registerPack packages.RegistrationPack
	bindErr := c.ShouldBindJSON(&registerPack)
	if bindErr != nil {
		a.DefaultBadRequest(c)
		return
	}
	user, err := a.AppContext.RegistrationService.Signup(registerPack)
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	a.Success(c, user)
}

func (a *AuthenticationController) Login(c *gin.Context) {
	var authenticationPack packages.AuthenticationPack
	bindErr := c.ShouldBindJSON(&authenticationPack)
	if bindErr != nil {
		a.DefaultBadRequest(c)
		return
	}
	authResponsePack, err := a.AppContext.AuthService.Authenticate(authenticationPack)
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	a.Success(c, authResponsePack)
}
