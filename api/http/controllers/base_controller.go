package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	http2 "vulscan/api/http"
	"vulscan/src/models"
)

type baseController struct {
	AppContext *http2.ApplicationContext
}

func (b *baseController) Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}

func (b *baseController) DefaultBadRequest(c *gin.Context) {
	c.JSON(http.StatusBadRequest, "Invalid request")
	c.Abort()
}

func (b *baseController) BadRequest(c *gin.Context, message string) {
	c.JSON(http.StatusBadRequest, message)
}

func (b *baseController) ErrorInternalServer(c *gin.Context) {
	c.JSON(http.StatusInternalServerError, "System error")
}

func (b *baseController) Unauthorized(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, "Unauthorized")
}

func (b *baseController) Forbidden(c *gin.Context) {
	c.JSON(http.StatusForbidden, "No permission to access this resource")
}

func (b *baseController) NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, "")
}

func (b *baseController) GetCurrentUser(c *gin.Context) *models.User {
	user, exist := c.Get("user")
	if !exist {
		b.Unauthorized(c)
		c.Abort()
		return nil
	}
	return user.(*models.User)
}
