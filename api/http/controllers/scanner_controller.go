package controllers

import (
	"github.com/gin-gonic/gin"
	"vulscan/api/http/context"
	"vulscan/src/packages"
)

type ScannerController struct {
	baseController
}

func NewScannerController(appContext *context.ApplicationContext) *ScannerController {
	return &ScannerController{
		baseController{
			AppContext: appContext,
		},
	}
}

func (s *ScannerController) Scan(c *gin.Context) {
	var scannerSegmentPack *packages.ScanSegmentPack
	if err := c.ShouldBindJSON(&scannerSegmentPack); err != nil {
		s.DefaultBadRequest(c)
		return
	}
	segment, err := s.AppContext.ScannerService.ScanSegment(scannerSegmentPack, s.GetCurrentUser(c))
	if err != nil {
		s.ErrorInternalServer(c)
		return
	}
	s.Success(c, segment)
}
