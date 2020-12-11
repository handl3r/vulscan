package controllers

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"vulscan/api/http"
)

type SegmentController struct {
	baseController
}

func NewSegmentController(appContext *http.ApplicationContext) *SegmentController {
	return &SegmentController{
		baseController{
			AppContext: appContext,
		},
	}
}

// Get controller get a segment by ID
func (s *SegmentController) Get(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		s.DefaultBadRequest(c)
		return
	}
	segment, err := s.AppContext.SegmentService.GetByID(id, s.GetCurrentUser(c))
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	responseData, jsonErr := json.Marshal(segment)
	if jsonErr != nil {
		s.ErrorInternalServer(c)
		return
	}
	s.Success(c, responseData)
}

// Delete controller delete a segment by ID
func (s *SegmentController) Delete(c *gin.Context) {
	id := c.Param("id")
	if len(id) == 0 {
		s.DefaultBadRequest(c)
		return
	}
	err := s.AppContext.SegmentService.DeleteByID(id, s.GetCurrentUser(c))
	if err != nil {
		c.JSON(err.GetHttpCode(), err.GetMessage())
		return
	}
	s.NoContent(c)
}

// TODO Add controller to receive targets from another segment to create new segment
