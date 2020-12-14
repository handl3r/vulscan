package controllers

import (
	"github.com/gin-gonic/gin"
	"vulscan/api/http/context"
)

type UserController struct {
	baseController
}

func NewUserController(appContext *context.ApplicationContext) *UserController {
	return &UserController{
		baseController{
			AppContext: appContext,
		},
	}
}

func (u *UserController) GetProjectsByUserID(c *gin.Context) {
	currentUser := u.GetCurrentUser(c)
	if currentUser == nil {
		u.Unauthorized(c)
		return
	}
	projects, err := u.AppContext.UserService.GetProjectByUser(currentUser)
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	u.Success(c, projects)
}
