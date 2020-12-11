package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"vulscan/api/http"
	"vulscan/src/packages"
)

type AuthenticationController struct {
	baseController
}

func NewAuthenticationController(appContext *http.ApplicationContext) *AuthenticationController {
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
	responseData, bindErr := json.Marshal(user)
	if bindErr != nil {
		a.ErrorInternalServer(c)
		return
	}
	a.Success(c, responseData)
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
	responseData, bindErr := json.Marshal(authResponsePack)
	if bindErr != nil {
		a.ErrorInternalServer(c)
		return
	}
	a.Success(c, responseData)
}
