package http

import (
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	Router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
	return &Server{
		Router: router,
	}
}

func (s *Server) Run(address string) {
	err := s.Router.Run(address)
	if err != nil {
		log.Fatalf("Error when start server: %s", err)
	}
	log.Printf("Server run on %s", address)
}
