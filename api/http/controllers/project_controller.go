package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"vulscan/src/packages"
)

type ProjectController struct {
	baseController
}

func NewProjectController(appContext *ApplicationContext) *ProjectController {
	return &ProjectController{
		baseController{
			AppContext: appContext,
		},
	}
}

// Get controller get project with full information by id
func (p *ProjectController) Get(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		p.DefaultBadRequest(c)
		return
	}
	p.AppContext.ProjectService.GetByID()

}

// Create controller create project
func (p *ProjectController) Create(c *gin.Context) {
	var createProjectPack *packages.CreateProjectPack
	if err := c.ShouldBindJSON(createProjectPack); err != nil {
		p.DefaultBadRequest(c)
		return
	}
	project, err := p.AppContext.ProjectService.Create(createProjectPack, p.GetCurrentUser(c))
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	responseData, newError := json.Marshal(project)
	if newError != nil {
		p.ErrorInternalServer(c)
		return
	}
	p.Success(c, responseData)
}

// Update controller update project
func (p *ProjectController) Update(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		p.DefaultBadRequest(c)
		return
	}
	var updateProjectPack packages.UpdateProjectPack
	if err := c.ShouldBindJSON(&updateProjectPack); err != nil {
		p.DefaultBadRequest(c)
		return
	}
	existProject, err := p.AppContext.ProjectService.GetByID(updateProjectPack.ID)
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	if existProject.User.ID != p.GetCurrentUser(c).ID {
		p.Unauthorized(c)
		return
	}
	updatedProject, err := p.AppContext.ProjectService.Update(&updateProjectPack)
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
	}
	dataResponse, jsonErr := json.Marshal(updatedProject)
	if jsonErr != nil {
		p.ErrorInternalServer(c)
		return
	}
	p.Success(c, dataResponse)
}

// Delete controller delete a project return http.StatusNoContent if success
func (p *ProjectController) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		p.DefaultBadRequest(c)
	}
	err := p.AppContext.ProjectService.Delete(id)
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	p.NoContent(c)
}
