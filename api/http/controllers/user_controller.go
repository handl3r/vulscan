package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	baseController
}

func NewUserController(appContext *ApplicationContext) *UserController {
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
	responseData, jsonErr := json.Marshal(projects)
	if jsonErr != nil {
		u.ErrorInternalServer(c)
		return
	}
	u.Success(c, responseData)
}
