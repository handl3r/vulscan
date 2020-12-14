package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"vulscan/api/http/context"
	"vulscan/src/packages"
)

type ProjectController struct {
	baseController
}

func NewProjectController(appContext *context.ApplicationContext) *ProjectController {
	return &ProjectController{
		baseController{
			AppContext: appContext,
		},
	}
}

// Get controller get project with all segments by id
func (p *ProjectController) Get(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		p.DefaultBadRequest(c)
		return
	}
	project, err := p.AppContext.ProjectService.GetByID(id, p.GetCurrentUser(c))
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	p.Success(c, project)
}

// Create controller create project
func (p *ProjectController) Create(c *gin.Context) {
	var createProjectPack *packages.CreateProjectPack
	if err := c.ShouldBindJSON(&createProjectPack); err != nil {
		log.Printf("Error when bindJSOn for creatProjectPack: %s", err)
		p.DefaultBadRequest(c)
		return
	}
	project, err := p.AppContext.ProjectService.Create(createProjectPack, p.GetCurrentUser(c))
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	p.Success(c, project)
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
	updatedProject, err := p.AppContext.ProjectService.Update(&updateProjectPack, p.GetCurrentUser(c))
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
	}
	p.Success(c, updatedProject)
}

// Delete controller delete a project return http.StatusNoContent if success
func (p *ProjectController) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		p.DefaultBadRequest(c)
	}
	err := p.AppContext.ProjectService.DeleteByID(id, p.GetCurrentUser(c))
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	p.NoContent(c)
}

func (p *ProjectController) Discover(c *gin.Context) {
	var discoverProjectPack *packages.DiscoverProjectPack
	if err := c.ShouldBindJSON(&discoverProjectPack); err != nil {
		p.DefaultBadRequest(c)
		return
	}
	segment, err := p.AppContext.ProjectService.Crawl(discoverProjectPack, p.GetCurrentUser(c))
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	p.Success(c, segment)
}
